package httpd

import (
	"coffeechat/api/cim"
	"coffeechat/pkg/logger"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"math/rand"
	"time"
)

// all grpc client connections
var logicClientMap map[int]*LogicGrpcClient = nil
var kMaxLogicClientConnect = 0
var serversCount = 0

const kMaxCheckInterval = 32 // s

type LogicGrpcClient struct {
	instance    cim.LogicClient
	conn        *grpc.ClientConn
	config      GrpcConfig
	index       int
	isConnected bool
}

func init() {
	logicClientMap = make(map[int]*LogicGrpcClient)
	c := time.Tick(time.Duration(time.Second * 1))

	go func() {
		var checkTime int64 = 1
		var tick int64 = 0
		for {
			<-c

			// 1s 2s 4s 8s 16s 32s
			if tick%checkTime == 0 {
				var isConnected = true
				for i := range logicClientMap {
					if logicClientMap[i].conn.GetState() != connectivity.Ready {
						logger.Sugar.Errorf("grpc server not connected,%s:%d,index=%d",
							logicClientMap[i].config.Ip, logicClientMap[i].config.Port, logicClientMap[i].index)
						isConnected = false
						logicClientMap[i].isConnected = false
					} else {
						if !logicClientMap[i].isConnected {
							// gRPC ping
							for j := 1; j <= 3; j++ {
								err := ping(logicClientMap[i].instance)
								if err != nil {
									logger.Sugar.Infof("grpc server connected success,but sayHello failed,retry=%d,%s:%d,index=%d,err=%s",
										j, logicClientMap[i].config.Ip, logicClientMap[i].config.Port, logicClientMap[i].index, err.Error())
									time.Sleep(10 * time.Millisecond)
								} else {
									logger.Sugar.Infof("grpc server connected success,%s:%d,index=%d",
										logicClientMap[i].config.Ip, logicClientMap[i].config.Port, logicClientMap[i].index)
									logicClientMap[i].isConnected = true
									break
								}
							}
						}
					}
				}
				if checkTime > kMaxCheckInterval {
					checkTime = 1
				}
				if !isConnected {
					checkTime *= 2
				}
			}

			tick++
		}
	}()
}

func StartGrpcClient(config []GrpcConfig) {
	logger.Sugar.Info("start grpc client")

	if len(config) < 2 {
		logger.Sugar.Fatal("at least 2 logic connections are required")
		return
	}
	serversCount = len(config)

	for i := range config {
		logger.Sugar.Infof("logic rpc server ip=%s,port=%d", config[i].Ip, config[i].Port)

		address := fmt.Sprintf("%s:%d", config[i].Ip, config[i].Port)
		//maxConnCnt := config[i].MaxConnCnt

		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			logger.Sugar.Error("connect grpc server=%s,error:", address, err.Error())
		} else {
			client := cim.NewLogicClient(conn)
			// save
			logicClientMap[kMaxLogicClientConnect] = &LogicGrpcClient{
				instance:    client,
				conn:        conn,
				config:      config[i],
				index:       i,
				isConnected: true,
			}
			// ping
			err := ping(client)
			if err != nil {
				// if failed, the routine will try again
				logicClientMap[kMaxLogicClientConnect].isConnected = false
			}
		}
		kMaxLogicClientConnect++
	}
}

func ping(conn cim.LogicClient) error {
	heart := &cim.CIMHeartBeat{}
	_, err := conn.Ping(context.Background(), heart)
	return err
}

// GetConn 获取随机一个可用的业务连接 0-1
func GetConn() cim.LogicClient {
	return logicClientMap[rand.Int()%serversCount].instance
}

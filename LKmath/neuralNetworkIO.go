package LKmath
//
//import (
//	"bufio"
//	"database/sql"
//	"encoding/json"
//	"fmt"
//	_ "github.com/go-sql-driver/mysql"
//	"io/ioutil"
//	"os"
//	"time"
//)
//
//
//type NeuralNetworkData struct {
//	database *sql.dB
//	err error
//}
//
//func NewNeuralNetworkData() NeuralNetworkData { //新建查询
//	data:= NeuralNetworkData{}
//	return data
//}
//
//
//func (this *NeuralNetworkData)ConnectToDatabase(driverName string, dataSourceName string)error{ //连接到数据库
//	this.database,this.err = sql.Open(driverName, dataSourceName)
//	if this.err!= nil{
//		fmt.Printf("%s",this.err)
//		return this.err
//
//	}
//	this.err = this.database.Ping()
//	if this.err != nil{
//		fmt.Printf("error in connecting database%s\n",this.err)
//		return this.err
//	}
//
//	fmt.Printf("database get connected successfully\n")
//	return nil
//}
//
//func (this *NeuralNetworkData)Close()error{ //关闭连接
//	this.err = this.database.Close()
//	if this.err != nil{
//		fmt.Printf("error occored when trying to close databse %s\n",this.err)
//		return this.err
//	}
//
//
//	return nil
//}
//
//
//
//
//
//func (this *NeuralNetworkData)Insert(NeuralNetworkName string, NeuralNetwork NeuralNetwork)error{
//
//	temp :=0.0
//	for i :=0; i < NeuralNetwork.HiddenLayerNum+1; i++{
//		for j :=0; j <NeuralNetwork.LayerParameter[i].NextLayerNum; j++{
//			for k :=-1; k<NeuralNetwork.LayerParameter[i].LayerParameter[j].W.Column;k++{
//
//				this.err = this.database.QueryRow("SELECT W_Value FROM neural_network.neural_network_values WHERE Neural_Network_Name = ? and Layer_No= ? and Node_No = ? and W_No = ?", NeuralNetworkName, i, j, k).Scan(&temp)
//
//				if this.err == sql.ErrNoRows{
//					if k==-1{
//						_, this.err = this.database.Exec("INSERT INTO neural_network.neural_network_values(Neural_Network_Name, Layer_No, Node_No, W_No, W_value) VALUES(?, ?, ?, ?, ?)",
//							NeuralNetworkName, i, j, k, NeuralNetwork.LayerParameter[i].LayerParameter[j].B)
//					}else{
//						_, this.err = this.database.Exec("INSERT INTO neural_network.neural_network_values(Neural_Network_Name, Layer_No, Node_No, W_No, W_value) VALUES(?, ?, ?, ?, ?)",
//							NeuralNetworkName, i, j, k, NeuralNetwork.LayerParameter[i].LayerParameter[j].W.Cell[0][k])
//					}
//					if this.err !=nil{
//						return this.err
//					}
//
//				}else if this.err == nil{
//					if k==-1{
//						_, this.err = this.database.Exec("UPDATE  neural_network.neural_network_values SET W_Value = ? WHERE Neural_Network_Name = ? and Layer_No= ? and Node_No = ? and W_No = ? ",
//							NeuralNetwork.LayerParameter[i].LayerParameter[j].B,NeuralNetworkName, i, j, k)
//					}else{
//						_, this.err = this.database.Exec("UPDATE  neural_network.neural_network_values SET W_Value = ? WHERE Neural_Network_Name = ? and Layer_No= ? and Node_No = ? and W_No = ? ",
//							NeuralNetwork.LayerParameter[i].LayerParameter[j].W.Cell[0][k],NeuralNetworkName, i, j, k)
//					}
//					if this.err !=nil{
//						fmt.Printf("%v",this.err)
//						return this.err
//					}
//				}
//
//
//
//			}
//
//		}
//	}
//
//	fmt.Printf("finishing inserting\n")
//	return nil
//
//
//}
//
//
//func (this *NeuralNetworkData)ReadFromDatabase(NeuralNetworkName string, neuralNetwork NeuralNetwork)(NeuralNetwork, error){
//
//	temp :=0.0
//
//	for i :=0; i < neuralNetwork.HiddenLayerNum+1; i++{
//		for j :=0; j < neuralNetwork.LayerParameter[i].NextLayerNum; j++{
//			for k :=-1; k< neuralNetwork.LayerParameter[i].LayerParameter[j].W.Column;k++{
//				this.err = this.database.QueryRow("SELECT W_Value FROM neural_network.neural_network_values WHERE Neural_Network_Name = ? and Layer_No= ? and Node_No = ? and W_No = ?", NeuralNetworkName, i, j, k).Scan(&temp)
//
//				if k==-1{
//					neuralNetwork.LayerParameter[i].LayerParameter[j].B = temp
//				}else{
//					neuralNetwork.LayerParameter[i].LayerParameter[j].W.Cell[0][k] = temp
//				}
//				if this.err !=nil{
//					fmt.Printf("an error occurred %v", this.err)
//					return NeuralNetwork{},this.err
//				}
//			}
//
//		}
//	}
//
//	return neuralNetwork, nil
//}
//
//func (this *NeuralNetworkData)Update(NeuralNetworkName string, neuralNetwork NeuralNetwork)error{
//
//	for i :=0; i < neuralNetwork.HiddenLayerNum+1; i++{
//		for j :=0; j < neuralNetwork.LayerParameter[i].NextLayerNum; j++{
//			for k :=-1; k< neuralNetwork.LayerParameter[i].LayerParameter[j].W.Column;k++{
//
//				if k==-1{
//					_, this.err = this.database.Exec("UPDATE  neural_network.neural_network_values SET W_Value = ? WHERE Neural_Network_Name = ? and Layer_No= ? and Node_No = ? and W_No = ? ",
//						neuralNetwork.LayerParameter[i].LayerParameter[j].B,NeuralNetworkName, i, j, k)
//				}else{
//					_, this.err = this.database.Exec("UPDATE  neural_network.neural_network_values SET W_Value = ? WHERE Neural_Network_Name = ? and Layer_No= ? and Node_No = ? and W_No = ? ",
//						neuralNetwork.LayerParameter[i].LayerParameter[j].W.Cell[0][k],NeuralNetworkName, i, j, k)
//				}
//
//				if this.err == sql.ErrNoRows{
//					if k==-1{
//						_, this.err = this.database.Exec("INSERT INTO neural_network.neural_network_values(Neural_Network_Name, Layer_No, Node_No, W_No, W_value) VALUES(?, ?, ?, ?, ?)",
//							NeuralNetworkName, i, j, k, neuralNetwork.LayerParameter[i].LayerParameter[j].B)
//					}else{
//						_, this.err = this.database.Exec("INSERT INTO neural_network.neural_network_values(Neural_Network_Name, Layer_No, Node_No, W_No, W_value) VALUES(?, ?, ?, ?, ?)",
//							NeuralNetworkName, i, j, k, neuralNetwork.LayerParameter[i].LayerParameter[j].W.Cell[0][k])
//					}
//
//				}else if this.err !=nil{
//					return this.err
//				}
//			}
//
//		}
//	}
//	return nil
//
//
//}
//
//
////to prevent error in writing
//func (this *NeuralNetwork)SaveToJson(dir string){
//
//	this.saveToJsonAsTemp(dir)
//
//	for ;;{
//		err:=os.Remove(dir +this.Name+".json")
//		if err!=nil{
//			fmt.Printf("error:%v",err)
//			time.Sleep(time.Second)
//			break
//		}else{
//			break
//		}
//
//	}
//
//	for ;;{
//		err:=os.Rename(dir+this.Name+"(temp).json", dir+this.Name+".json")
//		if err!=nil{
//			fmt.Printf("error:%v",err)
//			time.Sleep(time.Second)
//		}else{
//			break
//		}
//
//	}
//
//
//}
//
//const DefaultNeuralNetworkDirection ="data/neural_network_data/"
//
//func (this *NeuralNetwork) saveToJsonAsTemp(dir string){
//
//	JSON, MarshalErr := json.MarshalIndent(&this, "", "\t")
//	if MarshalErr !=nil{
//		panic(MarshalErr)
//	}
//
//	fileAddress := dir +this.Name+"(temp).json"
//	outputFile, outputError := os.OpenFile(fileAddress, os.O_CREATE, 0666)
//	if outputError != nil {
//		panic(outputError)
//	}
//	defer outputFile.Close()
//	outputWriter := bufio.NewWriter(outputFile)
//	outputWriter.WriteString(string(JSON))
//	outputWriter.Flush()
//
//	fmt.Printf("successfully save neuralnetwork to %s\n",fileAddress)
//}
//
//
//func (this *NeuralNetwork)ReadFromJson(dir string,neuralNetworkName string){
//
//	fileAddress := dir+neuralNetworkName+".json"
//	data, err := ioutil.ReadFile(fileAddress)
//	if err != nil {
//		panic(err)
//	} else {
//		fmt.Printf("Loading data from %s\n", fileAddress)
//	}
//
//	unmarshalErr := json.Unmarshal(data, &this)
//	if unmarshalErr != nil {
//		panic(unmarshalErr)
//	}
//
//	fmt.Printf("successfully read neural network from %s\n",fileAddress)
//
//}
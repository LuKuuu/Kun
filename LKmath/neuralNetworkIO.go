package LKmath

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)


type NeuralNetworkData struct {
	database *sql.DB
	err error
}

func NewNeuralNetworkData() NeuralNetworkData { //新建查询
	data:= NeuralNetworkData{}
	return data
}


func (this *NeuralNetworkData)ConnectToDatabase(driverName string, dataSourceName string)error{ //连接到数据库
	this.database,this.err = sql.Open(driverName, dataSourceName)
	if this.err!= nil{
		fmt.Printf("%s",this.err)
		return this.err

	}
	this.err = this.database.Ping()
	if this.err != nil{
		fmt.Printf("error in connecting database%s\n",this.err)
		return this.err
	}

	fmt.Printf("database get connected successfully\n")
	return nil
}

func (this *NeuralNetworkData)Close()error{ //关闭连接
	this.err = this.database.Close()
	if this.err != nil{
		fmt.Printf("error occored when trying to close databse %s\n",this.err)
		return this.err
	}
	return nil
}





func (this *NeuralNetworkData)Insert(NeuralNetworkName string, NeuralNetwork NeuralNetwork)error{

	temp :=0.0
	for i :=0; i < NeuralNetwork.HiddenLayerNum+1; i++{
		for j :=0; j <NeuralNetwork.LayerParameter[i].NextLayerNum; j++{
			for k :=-1; k<NeuralNetwork.LayerParameter[i].NodeParameter[j].W.Column;k++{

				this.err = this.database.QueryRow("SELECT W_Value FROM neural_network.neural_network_values WHERE Neural_Network_Name = ? and Layer_No= ? and Node_No = ? and W_No = ?", NeuralNetworkName, i, j, k).Scan(&temp)

				if this.err == sql.ErrNoRows{
					if k==-1{
						_, this.err = this.database.Exec("INSERT INTO neural_network.neural_network_values(Neural_Network_Name, Layer_No, Node_No, W_No, W_value) VALUES(?, ?, ?, ?, ?)",
							NeuralNetworkName, i, j, k, NeuralNetwork.LayerParameter[i].NodeParameter[j].B)
					}else{
						_, this.err = this.database.Exec("INSERT INTO neural_network.neural_network_values(Neural_Network_Name, Layer_No, Node_No, W_No, W_value) VALUES(?, ?, ?, ?, ?)",
							NeuralNetworkName, i, j, k, NeuralNetwork.LayerParameter[i].NodeParameter[j].W.Cell[0][k])
					}
					if this.err !=nil{
						return this.err
					}

				}else if this.err == nil{
					if k==-1{
						_, this.err = this.database.Exec("UPDATE  neural_network.neural_network_values SET W_Value = ? WHERE Neural_Network_Name = ? and Layer_No= ? and Node_No = ? and W_No = ? ",
							NeuralNetwork.LayerParameter[i].NodeParameter[j].B,NeuralNetworkName, i, j, k)
					}else{
						_, this.err = this.database.Exec("UPDATE  neural_network.neural_network_values SET W_Value = ? WHERE Neural_Network_Name = ? and Layer_No= ? and Node_No = ? and W_No = ? ",
							NeuralNetwork.LayerParameter[i].NodeParameter[j].W.Cell[0][k],NeuralNetworkName, i, j, k)
					}
					if this.err !=nil{
						fmt.Printf("%v",this.err)
						return this.err
					}
				}



			}

		}
	}
	return nil


}


func (this *NeuralNetworkData)ReadFromDatabase(NeuralNetworkName string, neuralNetwork NeuralNetwork)(NeuralNetwork, error){

	temp :=0.0

	for i :=0; i < neuralNetwork.HiddenLayerNum+1; i++{
		for j :=0; j < neuralNetwork.LayerParameter[i].NextLayerNum; j++{
			for k :=-1; k< neuralNetwork.LayerParameter[i].NodeParameter[j].W.Column;k++{
				this.err = this.database.QueryRow("SELECT W_Value FROM neural_network.neural_network_values WHERE Neural_Network_Name = ? and Layer_No= ? and Node_No = ? and W_No = ?", NeuralNetworkName, i, j, k).Scan(&temp)

				if k==-1{
					neuralNetwork.LayerParameter[i].NodeParameter[j].B = temp
				}else{
					neuralNetwork.LayerParameter[i].NodeParameter[j].W.Cell[0][k] = temp
				}
				if this.err !=nil{
					fmt.Printf("an error occurred %v", this.err)
					return NeuralNetwork{},this.err
				}
			}

		}
	}

	return neuralNetwork, nil
}

func (this *NeuralNetworkData)Update(NeuralNetworkName string, neuralNetwork NeuralNetwork)error{

	for i :=0; i < neuralNetwork.HiddenLayerNum+1; i++{
		for j :=0; j < neuralNetwork.LayerParameter[i].NextLayerNum; j++{
			for k :=-1; k< neuralNetwork.LayerParameter[i].NodeParameter[j].W.Column;k++{

				if k==-1{
					_, this.err = this.database.Exec("UPDATE  neural_network.neural_network_values SET W_Value = ? WHERE Neural_Network_Name = ? and Layer_No= ? and Node_No = ? and W_No = ? ",
						neuralNetwork.LayerParameter[i].NodeParameter[j].B,NeuralNetworkName, i, j, k)
				}else{
					_, this.err = this.database.Exec("UPDATE  neural_network.neural_network_values SET W_Value = ? WHERE Neural_Network_Name = ? and Layer_No= ? and Node_No = ? and W_No = ? ",
						neuralNetwork.LayerParameter[i].NodeParameter[j].W.Cell[0][k],NeuralNetworkName, i, j, k)
				}

				if this.err == sql.ErrNoRows{
					if k==-1{
						_, this.err = this.database.Exec("INSERT INTO neural_network.neural_network_values(Neural_Network_Name, Layer_No, Node_No, W_No, W_value) VALUES(?, ?, ?, ?, ?)",
							NeuralNetworkName, i, j, k, neuralNetwork.LayerParameter[i].NodeParameter[j].B)
					}else{
						_, this.err = this.database.Exec("INSERT INTO neural_network.neural_network_values(Neural_Network_Name, Layer_No, Node_No, W_No, W_value) VALUES(?, ?, ?, ?, ?)",
							NeuralNetworkName, i, j, k, neuralNetwork.LayerParameter[i].NodeParameter[j].W.Cell[0][k])
					}

				}else if this.err !=nil{
					return this.err
				}
			}

		}
	}
	return nil


}

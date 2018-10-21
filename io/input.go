package io

import (
//	"encoding/csv"
//	"fmt"
	"github.com/LuKuuu/Kun/LKmath"
//	"io"
//	"os"
//	"strconv"
)

func ReadMatrixFromFile(fileName string)LKmath.Matrix {

	if fileName == "" {
		panic("no file name")
	}

//fmt.Printf("start reading file from %s", fileName)
//	file, err := os.Open("./" + fileName)
//	if err != nil {
//		panic(fmt.Sprintf("there is an error %v when trying ro read %s.", err, fileName))
//	}
//	defer file.Close()
//	reader := csv.NewReader(file)
//
//	recordNum := 0
//	for {
//		record, err := reader.Read()
//		if err == io.EOF {
//			logs.L.Info("黑名单读取完成")
//			logs.L.Info("共从黑名单中读取%d条信息", recordNum)
//			break
//		} else if err != nil {
//			logs.L.Info("共从黑名单中读取%d条信息", recordNum)
//			logs.L.Warn("错误:%v, 停止黑名单的读取", err)
//			return
//		}
//
//		attachTimes, err := strconv.Atoi(record[1])
//		if err != nil {
//			logs.L.Warn("%s攻击次数有误%v, ", record[0], err)
//		} else {
//			//设置威胁等级
//			if attachTimes > configInfo.AttackTimes {
//				logs.L.Debug("ip%s 将被设为黑名单", record[0])
//				ipTable[ugetip(record[0])] = configInfo.MaxBlackLevel
//				recordNum++
//			} else if attachTimes >= 0 {
//
//				//判断该名单是否存在
//				if level, exist := ipTable[ugetip(record[0])]; exist{
//					//已存在的情况
//					if level == configInfo.WhiteLevel{
//						logs.L.Debug("ip%s 已在白名单中， 不做处理", record[0])
//
//					}else if level == configInfo.MaxBlackLevel{
//						logs.L.Debug("ip%s 已在黑名单中， 不做处理", record[0])
//
//
//					}else if level > configInfo.MediumBlackLevel{
//						logs.L.Debug("ip%s 已在灰名单中, 并且威胁等级大于将要设置的等级， 不做处理", record[0])
//
//					}else if level <= configInfo.MediumBlackLevel{
//						logs.L.Debug("ip%s 已在灰名单中, 并且威胁等级小于等于将要设置灰名单的等级， 设置为新的更高的等级", record[0])
//						ipTable[ugetip(record[0])] = configInfo.MediumBlackLevel
//
//					}
//				}else{
//					//不存在的情况
//					logs.L.Debug("ip%s 将被设为灰名单", record[0])
//					ipTable[ugetip(record[0])] = configInfo.MediumBlackLevel
//					recordNum++
//				}
//
//			} else {
//				logs.L.Warn("攻击次数应为一个自然数， 不记录该ip地址")
//
//			}
//		}

	return LKmath.Matrix{}
}


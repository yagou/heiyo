package conn

import (
	"crypto/md5"
	"errors"
	"fmt"
	"net"
	"strconv"
)

type hy_conn struct {
	conn_rows  map[string]net.Conn // tcp连接
	conn_count int                 // 当前连接数统计
	hash_tag   int                 // 连接数的哈希值
}

var hy_conn_cache *hy_conn

func NewHYConn() *hy_conn {
	if hy_conn_cache == nil {
		hy_conn_cache = &hy_conn{
			conn_rows: make(map[string]net.Conn),
		}
	}
	return hy_conn_cache
}

// 判断连接数量
func (hyc *hy_conn) Add_conn(conn net.Conn) (string, error) {
	tags := hyc.GetTag()
	if _, ok := hyc.conn_rows[tags]; ok {
		// 当前的tag已经存在，系统出错
		return "", errors.New("hy_conn.conn_rows tag exist")
	} else {
		hyc.conn_rows[tags] = conn
	}

	// 引用计数加1
	hyc.conn_count++
	return tags, nil
}

// 根据tag获取连接
func (hyc *hy_conn) Get_conn(tags string) (net.Conn, error) {
	if conn, ok := hyc.conn_rows[tags]; ok {
		return conn, nil
	} else {
		return nil, errors.New("conn not exist")
	}
}

// 获取所有的连接
func (hyc *hy_conn) Get_conn_all() map[string]net.Conn {
	return hyc.conn_rows
}

// 删除连接
func (hyc *hy_conn) Del_conn(tags string) {
	if conn, ok := hyc.conn_rows[tags]; ok {
		// 释放连接
		conn.Close()
		delete(hyc.conn_rows, tags)
		// 引用计数减1
		hyc.conn_count--
	}
}

// 获取tag
func (hyc *hy_conn) GetTag() string {
	tags := fmt.Sprintf("%x", md5.New().Sum([]byte(strconv.Itoa(hyc.hash_tag))))
	hyc.hash_tag++
	return tags
}

package main

import (
        "fmt"
        "io/ioutil"
        "os"
)

// StorageStrategy 存储策略接口
type StorageStrategy interface {
        Save(name string, data []byte) error
}

// fileStorage 保存到文件
type fileStorage struct{}

// Save 保存数据到文件
func (s *fileStorage) Save(name string, data []byte) error {
        return ioutil.WriteFile(name, data, os.ModePerm)
}

// encryptFileStorage 加密保存到文件
type encryptFileStorage struct{}

// Save 加密并保存数据到文件
func (s *encryptFileStorage) Save(name string, data []byte) error {
        encryptedData, err := encrypt(data)
        if err != nil {
                return err
        }
        return ioutil.WriteFile(name, encryptedData, os.ModePerm)
}

// encrypt 加密数据的实现
func encrypt(data []byte) ([]byte, error) {
        // 这里实现加密算法，简单示例：反转数据
        n := len(data)
        encrypted := make([]byte, n)
        for i := range data {
                encrypted[i] = data[n-1-i]
        }
        return encrypted, nil
}

// StorageStrategyFactory 工厂方法模式，用于创建存储策略
type StorageStrategyFactory struct{}

// NewStorageStrategy 根据类型创建存储策略
func (f *StorageStrategyFactory) NewStorageStrategy(t string) (StorageStrategy, error) {
        switch t {
        case "file":
                return &fileStorage{}, nil
        case "encrypt_file":
                return &encryptFileStorage{}, nil
        default:
                return nil, fmt.Errorf("not found StorageStrategy: %s", t)
        }
}

// 使用示例
func main() {
        factory := &StorageStrategyFactory{}

        // 创建普通文件存储策略
        fileStrategy, err := factory.NewStorageStrategy("file")
        if err != nil {
                fmt.Println("Error:", err)
                return
        }
        err = fileStrategy.Save("example.txt", []byte("Hello, World!"))
        if err != nil {
                fmt.Println("Error:", err)
                return
        }

        // 创建加密文件存储策略
        encryptStrategy, err := factory.NewStorageStrategy("encrypt_file")
        if err != nil {
                fmt.Println("Error:", err)
                return
        }
        err = encryptStrategy.Save("encrypted_example.txt", []byte("Hello, World!"))
        if err != nil {
                fmt.Println("Error:", err)
                return
        }

        fmt.Println("Files saved successfully!")
}

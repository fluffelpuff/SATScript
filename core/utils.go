package main

import "math/rand"

const letterBytes = "abcdefghklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ23"

func createNewConnectionId() string {
	b := make([]byte, 32)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func createNewPeerGroupId() string {
	b := make([]byte, 32)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

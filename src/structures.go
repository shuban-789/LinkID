package main

type block struct {
	Index        	int
	Initials     	string
	Sex			 	string
	Gender		 	string
	Age          	int
	Height       	float32
	Weight       	float32
	BMI				float32
	Blood		 	string
	Time         	string
	Location		string
	PreviousHash 	string
	CurrentHash  	string
	Prescriptions	[]string
	Conditions   	[]string
	VisitLogs		[]string
	History			[]string
}

type chain struct {
	ChainID    int
	BlockCount int
	Genesis    block
	Head       block
	Previous   block
	Chain      []block
}
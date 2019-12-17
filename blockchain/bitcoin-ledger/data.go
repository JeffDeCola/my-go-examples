// my-go-examples bitcoin-ledger data.go

package main

const (
	genesisTransactionString = `
        {
        "txID": 0,
        "inputs": [
                {
                    "inID": 0,
                    "refTxID": -1,
                    "inPubKey": "",
                    "signature": "Hello World, Welcome to the first transaction and block of jeffCoin"
                }
            ],
            "outputs": [
                {
                    "outPubKey": "` + foundersPubKey + `",
                    "value": 100000000
                }
            ]
        }`
	txRequestMessageSignedDataString1 = `
        {
            "txRequestMessage": {
                "sourceAddress": "` + foundersPubKey + `",
                "destinations": [
                    {
                        "destinationAddress": "` + jeffPubKey + `",
                        "value": 80000
                    }
                ]
            },
            "signature": "` + foundersSig + `"
        }`
	txRequestMessageSignedDataString2 = `
        {
            "txRequestMessage": {
                "sourceAddress": "` + jeffPubKey + `",
                "destinations": [
                    {
                        "destinationAddress": "` + mattPubKey + `",
                        "value": 50000
                    },
                    {
                        "destinationAddress": "` + coinVaultPubKey + `",
                        "value": 500
                    }
                ]
            },
            "signature": "` + jeffSig + `"
        }`
	txRequestMessageSignedDataString3 = `
        {
            "txRequestMessage": {
                "sourceAddress": "` + foundersPubKey + `",
                "destinations": [
                    {
                        "destinationAddress": "` + mattPubKey + `",
                        "value": 250000
                    }
                ]
            },
            "signature": "` + foundersSig + `"
        }`
	txRequestMessageSignedDataString4 = `
        {
            "txRequestMessage": {
                "sourceAddress": "` + mattPubKey + `",
                "destinations": [
                    {
                        "destinationAddress": "` + jillPubKey + `",
                        "value": 35000
                    }
                ]
            },
            "signature": "` + mattSig + `"
        }`
	txRequestMessageSignedDataString5 = `
        {
            "txRequestMessage": {
                "sourceAddress": "` + mattPubKey + `",
                "destinations": [
                    {
                        "destinationAddress": "` + jeffPubKey + `",
                        "value": 15000
                    }
                ]
            },
            "signature": "` + mattSig + `"
        }`
	txRequestMessageSignedDataStringBad = `
        {
            "txRequestMessage": {
                "sourceAddress": "` + foundersPubKey + `",
                "destinations": [
                    {
                        "destinationAddress": "Bad Pub Key",
                        "value": 80000
                    }
                ]
            },
            "signature": "Bad"
        }`
)

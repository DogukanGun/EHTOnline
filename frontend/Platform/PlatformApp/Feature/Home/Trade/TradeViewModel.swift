//
//  TradeViewModel.swift
//  PlatformApp
//
//  Created by Dogukan Gundogan on 21.10.23.
//

import Foundation

class TradeViewModel: ObservableObject {
    let options = [
        TradeOptions(symbol: "WETH", address: "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
        TradeOptions(symbol: "WBTC", address: "0x2260fac5e5542a773aa44fbcfedf7c193bc2c599"),
        TradeOptions(symbol: "LINK", address: "0x514910771af9ca656af840dff83e8264ecf986ca"),
        TradeOptions(symbol: "AAVE", address: "0x7Fc66500c84A76Ad7e9c93437bFc5Ac33E2DDaE9"),
        TradeOptions(symbol: "UNI", address: "0x1f9840a85d5af5bf1d1762f925bdaddc4201f984"),
        TradeOptions(symbol: "MATIC", address: "0x7D1AfA7B718fb893dB30A3aBc0Cfc608AaCfeBB0"),
        TradeOptions(symbol: "WBNB", address: "0x418D75f65a02b3D53B2418FB8E1fe493759c7605"),
        TradeOptions(symbol: "SNX", address: "0xc011a73ee8576fb46f5e1c5751ca3b9fe0af2a6f"),
        TradeOptions(symbol: "1INCH", address: "0x111111111117dc0aa78b770fa6a738034120c302"),
        TradeOptions(symbol: "CRV", address: "0xD533a949740bb3306d119CC777fa900bA034cd52"),
        TradeOptions(symbol: "SOL", address: "0xD533a949740bb3306d119CC777fa900bA034cd52"),
        TradeOptions(symbol: "stETH", address: "0xae7ab96520de3a18e5e111b5eaab095312d7fe84"),
        TradeOptions(symbol: "DAI", address: "0x6b175474e89094c44da98b954eedeac495271d0f"),
        TradeOptions(symbol: "USDT", address: "0xdac17f958d2ee523a2206206994597c13d831ec7")
    ]
    @Published var dexResponses:[DexResponse] = []
    @Published var selectedOption = ""
    
    func getDexData(tokenAddress:String) {
        networkManager.get(path: "/dex/" + tokenAddress, opts: nil) { (result:NetworkResponse<[DexResponse]>) in
            if let response = result.data {
                DispatchQueue.main.async {
                    self.dexResponses = response
                }
            }
        }
    }
}


struct TradeOptions:Hashable {
    var symbol:String
    var address:String
}

//
//  InvestmentViewModel.swift
//  PlatformApp
//
//  Created by Dogukan Gundogan on 22.10.23.
//

import Foundation
import SwiftUI
import metamask_ios_sdk
import Combine

@Observable class InvestmentViewModel {
    var selectedInvestmentType = ""
    var investmentOptions =  ["Withdraw","Deposit"]
    var supportedChains = ["Goerli","Binance"]
    var selectedDepositChain = ""
    var amount = ""
    var transactionStatus = ""
    private var ethereum = MetaMaskSDK.shared.ethereum
    private var cancellables: Set<AnyCancellable> = []
    
    func sendTransaction() {
        let parameters: [String: String] = [
            "to": "0x...", // receiver address
            "from": ethereum.selectedAddress, // sender address
            "value": amount
        ]
        
        let transactionRequest = EthereumRequest(
            method: .ethSendTransaction,
            params: [parameters]
        )
        
        ethereum.request(transactionRequest)?.sink(receiveCompletion: { completion in
            switch completion {
            case .failure(let error):
                print("\(error.localizedDescription)")
            default: break
            }
        }, receiveValue: { result in
            self.transactionStatus = (result as? String) ?? ""
        })
        .store(in: &cancellables)
    }
}

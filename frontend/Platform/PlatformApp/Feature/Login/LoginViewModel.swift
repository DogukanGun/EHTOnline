//
//  LoginViewModel.swift
//  PlatformApp
//
//  Created by Dogukan Gundogan on 18.10.23.
//

import Foundation
import SwiftUI
import metamask_ios_sdk
import Combine

class LoginViewModel: ObservableObject {
    private var cancellables: Set<AnyCancellable> = []
    private var ethereum = MetaMaskSDK.shared.ethereum
    private let dapp = Dapp(name: "Platform", url: "https://dogukangun.de")
    let accountRequest = EthereumRequest(method: .ethAccounts)
    var goToHomeScreen = false
    
    
    func login(completion:@escaping() -> Void){
        ethereum.connect(dapp)?.sink(receiveCompletion: { completion in
            switch completion {
            case let .failure(error):
                print("Connection error: \(error)")
            default: break
            }
        }, receiveValue: { result in
            UserDefaults.standard.setValue(result, forKey: "address")
            
        }).store(in: &cancellables)
        completion()
    }
}

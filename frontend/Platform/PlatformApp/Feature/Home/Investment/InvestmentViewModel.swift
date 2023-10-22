//
//  InvestmentViewModel.swift
//  PlatformApp
//
//  Created by Dogukan Gundogan on 22.10.23.
//

import Foundation
import SwiftUI

@Observable class InvestmentViewModel {
    var selectedInvestmentType = ""
    var investmentOptions =  ["Withdraw","Deposit"]
    var supportedChains = ["Goerli","Sepolia"]
    var selectedDepositChain = ""
    var amount = ""
}

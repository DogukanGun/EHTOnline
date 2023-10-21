//
//  OracleViewModel.swift
//  PlatformApp
//
//  Created by Dogukan Gundogan on 21.10.23.
//

import Foundation
import SwiftUI

@Observable class OracleViewModel {
    var oracleResponses:[OracleResponse] = []
    
    init() {
        getOraclePrices()
    }
    
    func getOraclePrices() {
        networkManager.get(path: "/prices", opts: nil) { (result:NetworkResponse<[OracleResponse]>) in
            if let data = result.data {
                self.oracleResponses = data
            }
        }
    }
}

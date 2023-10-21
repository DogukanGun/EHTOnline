//
//  OracleResponse.swift
//  PlatformApp
//
//  Created by Dogukan Gundogan on 21.10.23.
//

import Foundation

struct OracleResponse:Codable, Hashable {
    var price:Float
    var tokenDecimal:Int
    var oracleIdentifier:String
    var extraData:String
    
    enum CodingKeys: String,CodingKey {
        case price
        case tokenDecimal = "token_decimal"
        case oracleIdentifier = "oracle_identifier"
        case extraData = "extra_data"
    }
}

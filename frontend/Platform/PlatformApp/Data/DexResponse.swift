//
//  DexResponse.swift
//  PlatformApp
//
//  Created by Dogukan Gundogan on 21.10.23.
//

import Foundation

struct DexResponse: Codable, Hashable {
    var poolAssets:String
    var poolAddress:String
    var shouldSell:Bool
    var price:Float
    var advantage:Float
    
    enum CodingKeys: String,CodingKey {
        case poolAssets = "pool_assets"
        case poolAddress = "pool_address"
        case shouldSell = "should_sell"
        case price
        case advantage
    }
    
}

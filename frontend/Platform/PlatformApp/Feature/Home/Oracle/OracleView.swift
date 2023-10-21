//
//  OracleView.swift
//  PlatformApp
//
//  Created by Dogukan Gundogan on 19.10.23.
//

import SwiftUI

struct OracleView: View {
    @State var oracleViewModel = OracleViewModel()
    var body: some View {
        VStack {
            Text("OnChain Oracle Prices")
                .font(.largeTitle)
                .padding()
            List($oracleViewModel.oracleResponses,id: \.self) { oraclePrice in
                OracleViewListCell(oracleResponse: oraclePrice.wrappedValue)
            }
        }

    }
}

struct OracleViewListCell: View {
    var oracleResponse:OracleResponse
    var body: some View {
        HStack {
            VStack(alignment: .leading,spacing: 20) {
                Text("Source: \(oracleResponse.oracleIdentifier)")
                Text("\(oracleResponse.extraData) value: \(oracleResponse.price)")
            }
            Spacer()
            Image(systemName: "chevron.right")
        }
        .padding()
        .background(Color.white)
        .cornerRadius(10)
        .shadow(radius: 2)
        .padding()
    }
}

#Preview {
    OracleView()
}

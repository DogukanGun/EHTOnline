//
//  PriceView.swift
//  PlatformApp
//
//  Created by Dogukan Gundogan on 19.10.23.
//

import SwiftUI

struct TradeView: View {
    @ObservedObject var tradeViewModel = TradeViewModel()
    var body: some View {
        TradeViewPicker(options: tradeViewModel.options, selectedOption: $tradeViewModel.selectedOption) { token in
            self.tradeViewModel.getDexData(tokenAddress: token.address)
            self.tradeViewModel.selectedOption = token.symbol
        }.padding()
        List($tradeViewModel.dexResponses, id:\.self) { dex in
            TradeViewListCell(dexResponse: dex.wrappedValue,selectedCoin: self.tradeViewModel.selectedOption)
        }
    }
}

struct TradeViewPicker: View {
    var options:[TradeOptions]
    @Binding var selectedOption:String
    @State var isDropdownExpanded:Bool = false
    var onItemClick:(TradeOptions) -> Void
    var body: some View {
        VStack {
            DisclosureGroup("Select an option", isExpanded: $isDropdownExpanded) {
                ForEach(options, id: \.self) { option in
                    Button {
                        isDropdownExpanded = false
                        onItemClick(option)
                    } label: {
                        Text(option.symbol).tag(option.symbol)
                        
                    }
                }
            }
            .padding()
            .background(Color.white)
            .cornerRadius(10)
            .shadow(radius: 2)
            .onTapGesture {
                isDropdownExpanded.toggle()
            }
        }
    }
}

struct TradeViewListCell: View {
    var dexResponse:DexResponse
    var selectedCoin:String
    var body: some View {
        HStack {
            VStack(alignment: .leading,spacing: 20) {
                Text(dexResponse.poolAssets)
                Text("\(selectedCoin) value: \(dexResponse.price)")
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
    TradeView()
}

#Preview {
    TradeViewListCell(dexResponse: DexResponse(poolAssets: "ETH/BTC", poolAddress: "0x1esdqasd", shouldSell: false, price: 1540.2, advantage: 0.3232))
}


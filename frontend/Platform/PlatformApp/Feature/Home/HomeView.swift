//
//  Home.swift
//  PlatformApp
//
//  Created by Dogukan Gundogan on 19.10.23.
//

import SwiftUI

struct HomeView: View {
    var body: some View {
        TabView {
            OracleView()
                .tabItem {
                    HStack {
                        Image(systemName: "banknote")
                        Text("Oracle")
                    }
                }
            TradeView()
                .tabItem {
                    HStack {
                        Image(systemName: "dollarsign.arrow.circlepath")
                        Text("Trade")
                    }
                }
            InvestmentView()
                .tabItem {
                    HStack {
                        Image(systemName: "bitcoinsign.circle")
                        Text("Trade")
                    }
                }
        }
    }
}

#Preview {
    HomeView()
}

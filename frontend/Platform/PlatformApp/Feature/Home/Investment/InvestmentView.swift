//
//  InvestmentView.swift
//  PlatformApp
//
//  Created by Dogukan Gundogan on 21.10.23.
//

import SwiftUI

struct InvestmentView: View {
    @State var investmentViewModel = InvestmentViewModel()
    var body: some View {
        VStack {
            InvestmentTypeChooser(options: $investmentViewModel.investmentOptions, selectedOption: $investmentViewModel.selectedInvestmentType) { investmentType in
                self.investmentViewModel.selectedInvestmentType = investmentType
            }
            
            if investmentViewModel.selectedInvestmentType == "Deposit" {
                InvestmentTypeChooser(options: $investmentViewModel.supportedChains, selectedOption: $investmentViewModel.selectedDepositChain) { depositChain in
                    self.investmentViewModel.selectedDepositChain = depositChain
                }
            }
                        
            
            if investmentViewModel.selectedInvestmentType != "" {
                TextField(text: $investmentViewModel.amount) {
                    Text("Amount that you want to \(investmentViewModel.selectedInvestmentType.lowercased())")
                        .padding()
                        .foregroundColor(.black)
                }
                .padding()
                .background(Color.white)
                .cornerRadius(10)
                .shadow(radius: 2)
                .padding()
                Spacer()
                if investmentViewModel.transactionStatus != "" {
                    Text(investmentViewModel.transactionStatus)
                }
                Button {
                    investmentViewModel.sendTransaction()
                } label: {
                    Text(investmentViewModel.selectedInvestmentType)
                        .font(.custom("", size: 20))
                        .frame(maxWidth: .infinity)
                }
                .buttonStyle(BlueButton())
                .padding()
            }
            
        }
        
        
    }
}

struct InvestmentTypeChooser: View {
    @Binding var options:[String]
    @Binding var selectedOption:String
    @State var isDropdownExpanded:Bool = false
    var onItemClick:(String) -> Void
    var body: some View {
        VStack {
            DisclosureGroup(selectedOption == "" ? "Select an option" : selectedOption, isExpanded: $isDropdownExpanded) {
                ForEach(options, id: \.self) { option in
                    Button {
                        isDropdownExpanded = false
                        onItemClick(option)
                    } label: {
                        Text(option).tag(option)
                    }
                    .padding(.vertical)
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
        .padding()
    }
}

#Preview {
    InvestmentView()
}

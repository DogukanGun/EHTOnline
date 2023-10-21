//
//  Login.swift
//  PlatformApp
//
//  Created by Dogukan Gundogan on 18.10.23.
//

import SwiftUI

struct LoginView: View {
    @State private var navigationPath: [String] = []
    @ObservedObject var loginViewModel = LoginViewModel()
    var body: some View {
        NavigationStack(path: $navigationPath) {
            VStack(spacing:50) {
                LottieView(lottieFile: "WelcomeAnimation")
                    .frame(height: 200)
                LoginViaGmailButton{
                    loginViewModel.login(completion: {
                        navigationPath.append("home")
                    })
                }
            }
            .navigationDestination(for: String.self) { destination in
                if destination == "home" {
                    HomeView()
                        .navigationBarBackButtonHidden()
                } else {}
            }
        }
        
    }
}

struct LoginViaGmailButton: View {
    var buttonOnClick:()->Void
    var body: some View {
        Button {
            buttonOnClick()
        } label: {
            HStack(spacing: 30) {
                Image("metamask_logo")
                    .resizable()
                    .frame(width: 30, height: 30)
                Text("Login via Metamask")
                    .font(.title3)
            }
            
        }
        .buttonStyle(BlueButton())
    }
}


#Preview {
    LoginView()
}

//
//  ButtonDecoration.swift
//  PlatformApp
//
//  Created by Dogukan Gundogan on 18.10.23.
//

import Foundation
import SwiftUI

struct BlueButton: ButtonStyle {
    func makeBody(configuration: Configuration) -> some View {
        configuration.label
            .padding()
            .background(Color(red: 0, green: 0, blue: 0.4))
            .foregroundStyle(.white)
            .clipShape(Capsule())
    }
}

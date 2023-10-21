//
//  PlatformAppApp.swift
//  PlatformApp
//
//  Created by Dogukan Gundogan on 18.10.23.
//

import SwiftUI
import OSLog

let logger = Logger()

@main
struct PlatformAppApp: App {
    let persistenceController = PersistenceController.shared

    var body: some Scene {
        WindowGroup {
            LoginView()
                .environment(\.managedObjectContext, persistenceController.container.viewContext)
        }
    }
}

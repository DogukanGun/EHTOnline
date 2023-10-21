//
//  Network Manager.swift
//  PlatformApp
//
//  Created by Dogukan Gundogan on 21.10.23.
//

import OSLog

let sharedNetworkManager = NetworkManager()


class NetworkManager {
    var baseUrl = "http://127.0.0.1:3000"
    var networkManagerStatus:NetworkManagerStatus = .readyWithoutAuth
    let logger = Logger()

    
    func post<T:Codable, V:Encodable>(urlAsString:String, body:V, completion:@escaping(NetworkResponse<T>) -> Void) {
        guard let _url = URL(string: baseUrl + urlAsString) else {
            logger.error("URL cannot be parsed")
            return
        }
        var request = URLRequest(url:_url)
        request.httpMethod = "POST"
        do {
            let encodedBody = try JSONEncoder().encode(body)
            request.httpBody = encodedBody
            request.setValue("application/json", forHTTPHeaderField: "Content-Type")
        } catch {
            logger.error("An error occurred: \(error)")
            return
        }
        sendRequest(request: request, completion: completion)
    }
    
    func get<T:Codable>(path:String, opts:String?, completion:@escaping(NetworkResponse<T>) -> Void) {
        guard let _url = URL(string: baseUrl + path) else {
            logger.error("URL cannot be parsed")
            return
        }
        var request = URLRequest(url:_url)
        request.httpMethod = "GET"
        sendRequest(request: request, completion: completion)
    }
    
    func sendRequest<T:Codable>(request:URLRequest, completion:@escaping(NetworkResponse<T>) -> Void) {
        var networkResponse = NetworkResponse<T>()
        let task = URLSession.shared.dataTask(with: request) { data, response, error in
            if let error = error {
                self.logger.error("An error occurred: \(error)")
                completion(networkResponse)
                return
            }
            
            guard let httpResponse = response as? HTTPURLResponse else {
                self.logger.error("An error occurred in api: \(String(describing: error))")
                completion(networkResponse)
                return
            }
            
            guard (200..<300).contains(httpResponse.statusCode) else {
                self.logger.error("An error occurred in api: \(String(describing: error))")
                completion(networkResponse)
                return
            }
            
            guard let data = data else {
                self.logger.error("No data")
                completion(networkResponse)
                return
            }
            
            do {
                networkResponse = try JSONDecoder().decode(NetworkResponse<T>.self, from: data)
                completion(networkResponse)
            } catch {
                self.logger.error("An error occurred in response parsing: \(error)")
                completion(networkResponse)
            }
        }
        task.resume()
    }
}

class NetworkResponse<T:Codable>:Codable {
    var data:T?
    
    init(data: T? = nil) {
        self.data = data
    }
}

enum NetworkManagerStatus {
    case readyWithoutAuth
    case fail
    case readyWithAuth
}

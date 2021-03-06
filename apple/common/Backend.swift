import Foundation

class Backend {

    static let shared = Backend()
    private let network = Network()
    private var sessionId: String?
    private var crypto: Crypto?

    private enum Key: String {
        case username
        case password
    }

    struct Credential {
        let username: String
        let password: String
    }
    var credential: Credential?

    private init() {

//        UserDefaults.standard.set(nil, forKey: Key.username.rawValue)
//        UserDefaults.standard.set(nil, forKey: Key.password.rawValue)

        if let username = UserDefaults.standard.string(forKey: Key.username.rawValue),
            let password = UserDefaults.standard.string(forKey: Key.password.rawValue) {
            self.credential = Credential(username: username, password: password)
        }
    }

    private func send(_ haberBuilder:Haber.Builder) {
        guard let haber = try? haberBuilder.setSessionId(sessionId ?? "").build() else {
            print("could not create haber")
            return
        }
        print("write \(haber.data().count) bytes for \(haber.which)")
        self.network.send(haber.data())
    }

    func connect() {
        network.connect()
    }

    func sendText(_ body: String, to: String) {
        guard let update = try? Text.Builder().setBody(body).build() else {
            print("could not create Text")
            return
        }
        let haberBuilder = Haber.Builder().setText(update).setWhich(.text).setTo(to)
        send(haberBuilder)
    }

    func sendContacts(_ contacts: [String:Contact]) {
        let haberBuilder = Haber.Builder().setContacts(Array(contacts.values)).setWhich(.contacts)
        send(haberBuilder)
    }

    func sendStore(key: Data, value: Data) {
        do {
            guard let encrypted = crypto?.encrypt(data: value) else {
                print("could not encrypt store")
                return
            }
            let store = try Store.Builder().setKey(key).setValue(encrypted).build()
            let haberBuilder = Haber.Builder().setStore(store).setWhich(.store)
            send(haberBuilder)
        } catch {
            print(error.localizedDescription)
        }
    }

    private func didReceiveStore(_ haber: Haber) {
        guard let value = crypto?.decrypt(ciphertext: haber.store.value) else {
            print("could not decrypt store")
            return
        }
        Model.shared.didReceiveStore(key: haber.store.key, value: value)
    }

    func sendLoad(key: Data) {
        do {
            let load = try Load.Builder().setKey(key).build()
            let haberBuilder = Haber.Builder().setWhich(.load).setLoad(load)
            send(haberBuilder)
        } catch {
            print(error.localizedDescription)
        }
    }

    func register(username: String, password: String) {
        print("register not implemented")
    }

    func login(username: String, password: String) {
        do {
            credential = Credential(username: username, password: password)
            let login = try Login.Builder().setUsername(username).build()
            let haberBuilder = Haber.Builder().setLogin(login).setWhich(.login)
            send(haberBuilder)
        } catch {
            print(error.localizedDescription)
        }
    }

    func didReceiveData(_ data: Data) {
        guard let haber = try? Haber.parseFrom(data:data) else {
                print("Could not deserialize")
                return
        }
        if let sid = haber.sessionId, sessionId == nil {
            authenticated(sessionId: sid)
        }
    
        print("read \(data.count) bytes for \(haber.which)")
        switch haber.which {
            case .contacts: Model.shared.didReceiveContacts(haber.contacts)
            case .text:     Model.shared.didReceiveText(haber, data: data)
            case .presence: Model.shared.didReceivePresence(haber)
            case .store:    didReceiveStore(haber)
            default:        print("did not handle \(haber.which)")
        }
    }

    private func authenticated(sessionId sid: String) {
        sessionId = sid

        guard let username = self.credential?.username, let password = credential?.password else {
            print("authentication without credentials")
            return
        }
        UserDefaults.standard.set(username, forKey: Key.username.rawValue)
        UserDefaults.standard.set(password, forKey: Key.password.rawValue)
        crypto = Crypto(password: password)

        EventBus.post(.authenticated)
    }
}

VueWS = (function() {

const VueWS = {
        install(Vue, connectionUrl) {
                const wsConnection = WsConnection(connectionUrl)
                Vue.prototype.$wsocket = wsConnection.wsocket

                Vue.mixin({
                        created,
                        beforeDestroy
                })

                function created() {
                        wsConnection.listen(this)
                }
                
                function beforeDestroy() {
                        wsConnection.unlisten(this)
                }
        }
}

return VueWS

function WsConnection(url) {
        let isConnected = false
        const listenedComponents = {}

        const ws = new WebSocket(url)
        ws.onopen = function() {
                isConnected = true

                for (component of Object.values(listenedComponents)) {
                        const connected = component.$options.webSockets.connected;
                        if (null != connected) {
                                connected.call(component)
                        }
                }
        }
        ws.onerror = function() {
                isConnected = false
        }
        ws.onclose = function() {
                isConnected = false
        }
        ws.onmessage = function(event) {

        }

        function sendMessage(message) {
                if (!isConnected) {
                        throw new Error("Ws not ready")
                }
                ws.send(message)
        }

        const wsocket = {
                send: sendMessage
        }

        function listen(component) {
                const wsOptions = component.$options.webSockets
                if (null == wsOptions) {
                        return
                }
                listenedComponents[component._uid] = component
                
                if (isConnected && wsOptions.connected) {
                        wsOptions.connected.call(component)
                }
        }

        function unlisten(component) {
                delete listenedComponents[component._uid]
        }

        return {
                wsocket,
                listen,
                unlisten
        }
}

})()
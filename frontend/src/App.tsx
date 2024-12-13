import {Centrifuge} from "centrifuge";
import {useEffect} from "react";

export function App() {
    useEffect(() => {
        const centrifugo = new Centrifuge('ws://localhost:8000/connection/websocket', {
            getToken: async (ctx) => {
                let result = await fetch('http://localhost:8001/auth/jwt')
                let token = await result.text()

                console.log('Token:', token)

                return token;
            }
        });

        centrifugo.on('connecting', () => {
            console.log('Connecting to Centrifugo');
        });

        centrifugo.on('connected', () => {
            console.log('Connected to Centrifugo');
        });

        centrifugo.on('disconnected', () => {
            console.log('Disconnected from Centrifugo');
        });

        centrifugo.on('publication', (message) => {
            console.log('Publication:', message);
        })

        centrifugo.connect();
    }, [])
    return (
        <div>
            <h1>Centrifugo Tutorial</h1>
            <p>Open the browser console to see the output</p>
        </div>
    )
}
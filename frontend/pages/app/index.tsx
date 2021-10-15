import React, { Component } from 'react'
import Router from 'next/router'
import cookie from 'js-cookie'

const apiServer: string = "https://0ce60ffa-33aa-4d79-a622-8df6ffc00b3f.mock.pstmn.io";

interface Props {
    
}
interface State {
    apiData: any
}


export default class index extends Component<Props, State> {
    state = {
        apiData: "server hasn't returned any data"
    }

    componentDidMount() {
        const authCookie = cookie.get("auth")
        
        if (authCookie == undefined) {
            Router.push('/login')
        }

        const options: RequestInit = {
            method: 'POST',
            body: JSON.stringify({ 'Auth': authCookie }),
            headers: {
                'Content-Type': 'application/json'
            }
        }

        fetch(apiServer + '/api/auth/v1/session', options).then(e => {
            e.text().then(ee => {
                this.setState({
                    ...this.state,
                    apiData: ee
                })
            })
        })
    }

    render() {
        return (
            <div>
                {this.state.apiData}
            </div>
        )
    }
}

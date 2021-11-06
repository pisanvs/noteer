import React, { Component } from 'react'
import Router from 'next/router'
import cookie from 'js-cookie'

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
        const authCookie = cookie.get("sesh")
        
        if (authCookie == undefined) {
            Router.push('/login')
        }

        const options: RequestInit = {
            method: 'GET',
            headers: {
                'Authorization': authCookie!
            }
        }

        fetch('https://api.noteer.local/api/v1/auth/session', options).then(e => {
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

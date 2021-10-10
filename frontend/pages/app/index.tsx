import React, { Component } from 'react'
import Head from "next/head"

const apiServer: string = "https://0ce60ffa-33aa-4d79-a622-8df6ffc00b3f.mock.pstmn.io" as string;

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
        fetch(apiServer + '/data').then(e => {
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

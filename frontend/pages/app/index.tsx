/**
* Copyright (C) 2021  Maximiliano Morel (pisanvs) <maxmorel@pisanvs.cl>
*
* This file is part of Noteer, a note taking application.
* 
* Noteer is free software: you can redistribute it and/or modify
* it under the terms of the GNU General Public License v3 as
* published by the Free Software Foundation
*
* Noteer is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
* GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License
* along with Noteer.  If not, see <https://www.gnu.org/licenses/>.
*
*
* @license GPL-3.0 <https://www.gnu.org/licenses/gpl-3.0.txt>
*/

import React, { Component } from 'react'
import '../../styles/app.module.css'
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

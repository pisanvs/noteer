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

import Cookies from 'js-cookie';
import Head from 'next/head'
import React, { Component } from 'react'
import { diff_match_patch } from 'diff-match-patch'
import Showdown from 'showdown'

import styles from '../styles/app.module.css'

interface Props {
    
}
interface State {
    
}

export default class Editor extends Component<Props, State> {
    editor: React.RefObject<HTMLDivElement>;
    pastEditorContent: string | undefined;
    id: React.RefObject<HTMLInputElement>;

    constructor(props: Props) {
        super(props)
        this.state = {}

        this.editor = React.createRef();
        this.id = React.createRef();
    }

    componentDidMount() {
        const options: RequestInit = {
            method: 'POST',
            headers: {
                Authorization: Cookies.get('sesh')!,
                'Content-Type': 'application/x-www-form-urlencoded'
            },
            body: "id=619565bad3743043c4a57b18"
        }

        fetch("//api.noteer.local/api/v1/data/docs/get", options).then(d => {
            d.text().then(e => {
                this.pastEditorContent = e;
                this.editor.current!.innerHTML = e;
            })
        })
        this.pastEditorContent = this.editor.current?.innerHTML;
        setInterval(() => {
            if (this.editor.current?.innerHTML == this.pastEditorContent) {
                return;
            } else {
                this.sendEdits();
            }
        }, 5000)
    }

    setId = () => {
        const options: RequestInit = {
            method: 'POST',
            headers: {
                Authorization: Cookies.get('sesh')!,
                'Content-Type': 'application/x-www-form-urlencoded'
            },
            body: "id="+this.id.current!.value
        }

        fetch("//api.noteer.local/api/v1/data/docs/get", options).then(d => {
            d.text().then(e => {
                this.pastEditorContent = e;
                this.editor.current!.innerHTML = e;
            })
        })
    }

    sendEdits = () => {
        let dmp = new diff_match_patch()
        let prediff = dmp.diff_main(this.pastEditorContent!, this.editor.current!.innerHTML)

        let diff = dmp.patch_toText(dmp.patch_make(prediff))

        const options: RequestInit = {
            method: 'POST',
            headers: {
                Authorization: Cookies.get('sesh')!,
                'Content-Type': 'application/x-www-form-urlencoded'
            },
            body: "id=619565bad3743043c4a57b18&content="+encodeURIComponent(diff)
        }

        fetch("//api.noteer.local/api/v1/data/docs/edit", options).then(d => {
            if (d.status == 200) {
                this.pastEditorContent = this.editor.current!.innerHTML;
            } else {
                this.sendEdits();
            }
        })
    }

    render() {
        return (
            <>
                <Head>
                    <title>Noteer - Editor</title>
                </Head>

                <div className="content">
                    <input type="text" name="id" id="id" ref={this.id} />
                    <button onClick={this.setId}>Submit</button>
                    <h1>Editor</h1>
                    <div contentEditable ref={this.editor} className={styles.editor}></div>
                    <button onClick={this.sendEdits}>submit changes</button>
                </div>
            </>
        )
    }
}

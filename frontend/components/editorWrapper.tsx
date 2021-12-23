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

import Cookies from "js-cookie";
import Head from "next/head";
import React, { Component } from "react";

import { diff_match_patch } from "diff-match-patch";

import Editor from './editor.jsx'

interface Props {}
interface State {
  text: string;
  curPos: number;
}

export default class EditorWrapper extends Component<Props, State> {
  pastEditorContent: string | undefined;
  id: React.RefObject<HTMLInputElement>;
  editor: React.RefObject<HTMLDivElement>;

  constructor(props: Props) {
    super(props);
    this.state = {
      text: "",
      curPos: 0,
    };

    this.id = React.createRef();
    this.editor = React.createRef();
  }

  componentDidMount() {
    const options: RequestInit = {
      method: "POST",
      headers: {
        Authorization: Cookies.get("sesh")!,
        "Content-Type": "application/x-www-form-urlencoded",
      },
      body: "id=61ad11bc0b162e3d513fd26e",
    };

    fetch("//api.noteer.local/api/v1/data/docs/get", options).then((d) => {
      d.text().then((e) => {
        this.pastEditorContent = e;
        this.setState({ text: e });
      });
    });
    this.pastEditorContent = this.state.text;
    setInterval(() => {
      if (this.state.text == this.pastEditorContent) {
        return;
      } else {
        this.sendEdits();
      }
    }, 5000);
  }

  setId = () => {
    const options: RequestInit = {
      method: "POST",
      headers: {
        Authorization: Cookies.get("sesh")!,
        "Content-Type": "application/x-www-form-urlencoded",
      },
      body: "id=" + this.id.current!.value,
    };

    fetch("//api.noteer.local/api/v1/data/docs/get", options).then((d) => {
      d.text().then((e) => {
        this.pastEditorContent = e;
        this.setState({ text: e });
      });
    });
  };

  sendEdits = () => {
    let dmp = new diff_match_patch();
    let prediff = dmp.diff_main(this.pastEditorContent!, this.state.text);

    let diff = dmp.patch_toText(dmp.patch_make(prediff));

    const options: RequestInit = {
      method: "POST",
      headers: {
        Authorization: Cookies.get("sesh")!,
        "Content-Type": "application/x-www-form-urlencoded",
      },
      body:
        "id=" + this.id.current?.value + "&content=" + encodeURIComponent(diff),
    };

    fetch("//api.noteer.local/api/v1/data/docs/edit", options).then((d) => {
      if (d.status == 200) {
        this.pastEditorContent = this.state.text;
      } else {
        this.sendEdits();
      }
    });
  };

  updateText = () => {
    this.setState({ text: this.editor.current?.innerHTML! });
  };

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
          <Editor text={[{type: "paragraph", children: [{text: "Start Writing..."}]}]} />
          <button onClick={this.sendEdits}>submit changes</button>
        </div>
      </>
    );
  }
}

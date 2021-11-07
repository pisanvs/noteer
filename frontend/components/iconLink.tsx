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

import React from "react"
import styles from "../styles/index.module.css"

interface Props {
    text: string;
    href?: string;
    style?: React.CSSProperties;
}

export default function iconLink (props: React.PropsWithChildren<Props>) {
    return (
        <div className={styles["icon-link"]} style={props.style}>
            {props.children}{' '}<a className={styles["nav-boldlink"]} href={props.href}>{props.text}</a>
        </div>
    )
}

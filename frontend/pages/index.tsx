/* eslint-disable @next/next/no-img-element */
import React, { Component } from "react";

import styles from "../styles/index.module.css";

import { AiOutlineUser } from "@react-icons/all-files/ai/AiOutlineUser";
import { IoHelp } from "@react-icons/all-files/io5/IoHelp";
import { AiOutlineArrowUp } from "@react-icons/all-files/ai/AiOutlineArrowUp";

import Head from "next/head";
import Link from "next/link";

import IconLink from "../components/iconLink";
import Typed from "typed.js";

interface Props {}
interface State {}

export default class index extends Component<Props, State> {
  state = {};
  typed: Typed | undefined;

  componentDidMount() {
    this.typed = new Typed("#type-b", {
      strings: ["better", "smarter", "faster", "more private", "more powerful"],
      typeSpeed: 50,
      backSpeed: 50,
      loop: true,
      cursorChar: "&#9647;",
    });
  }

  render() {
    return (
      <>
        <Head>
          <title>Noteer, notes but better</title>
        </Head>
        <div className={styles["container"]}>
          <div className={styles["header"]}>
            <img src="/logo.svg" alt="logo" className={styles.logo}></img>
            <div className={styles["header-left"]}>
              <div className={styles.navlinks}>
                <Link href="/login" passHref>
                  <IconLink text="Login">
                    <AiOutlineUser className={styles["nav-icon"]} />
                  </IconLink>
                </Link>
                <Link href="/help" passHref>
                  <IconLink text="Help">
                    <IoHelp className={styles["nav-icon"]} />
                  </IconLink>
                </Link>
                <Link href="/app" passHref>
                  <IconLink text="Access">
                    <AiOutlineArrowUp className={styles["nav-icon"]} />
                  </IconLink>
                </Link>
              </div>
            </div>
            <div className={styles.filler} />
            <div className={styles["header-right"]}>
              <div className={styles.navlinks}>
                <Link href="/register">
                  <a className={styles["nav-thinlink"]}>Register</a>
                </Link>
                <Link href="/#about" passHref>
                  <a className={styles["nav-thinlink"]}>About Us</a>
                </Link>
                <Link href="https://github.com/pisanvs/noteer">
                  <a className={styles["nav-thinlink"]}>Github Repo</a>
                </Link>
              </div>
            </div>
          </div>
          <div className={styles["content"]}>
            <div className={styles["content-top"]}>
              <div className={styles["content-left"]}>
                <img
                  src="/devices.png"
                  alt="devices"
                  className={styles["devices-image"]}
                ></img>
              </div>

              <div className={styles["content-right"]}>
                <div className="content-right-top">
                  <p className={styles["title"]}>Notes but, </p>
                  <code className={styles["code"]}><span id="type-b"></span></code>
                </div>
                <div className={styles["content-right-bottom"]}>
                  <div className={styles["benefit-right"]}>
                    <span className={styles["text-bold"]}>Better</span>
                    <p className={styles["text-normal"]}>
                      Take better notes, and take your memory and knowledge to
                      another level. Decreasing the gap between ideas and words.
                    </p>
                  </div>
                  <div className={styles["benefit-left"]}>
                    <span className={styles["text-bold"]}>Faster</span>
                    <p className={styles["text-normal"]}>
                      The less time an idea takes from being in your mind to
                      being stored in a computer, the less chances it gets for
                      being forgotten.
                    </p>
                  </div>
                  <div className={styles["benefit-right"]}>
                    <span className={styles["text-bold"]}>Smarter</span>
                    <p className={styles["text-normal"]}>
                      Including Artificial Intelligence for note browsing,
                      completion, and summarizing. Noteer makes you a better
                      note-taker without making much effort.
                    </p>
                  </div>
                  <div className={styles["tryit"]}>
                    <button className={styles["primary-btn"]}>
                      Try it! It&lsquo;s free!
                    </button>
                    <div className={styles["tryit-sh"]}>
                      (and self hosted if you wish)
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div className="content-bottom">
              <h2>Make your notes useful</h2>
            </div>
          </div>
        </div>
      </>
    );
  }
}

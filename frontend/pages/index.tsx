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

/* eslint-disable @next/next/no-img-element */
import React, { Component } from "react";

import styles from "../styles/index.module.css";

import { AiOutlineUser } from "@react-icons/all-files/ai/AiOutlineUser";
import { IoHelp } from "@react-icons/all-files/io5/IoHelp";
import { AiOutlineArrowUp } from "@react-icons/all-files/ai/AiOutlineArrowUp";
import { BsMoon } from "@react-icons/all-files/bs/BsMoon";
import { FiSun } from "@react-icons/all-files/fi/FiSun";

import Head from "next/head";
import Link from "next/link";

import IconLink from "../components/iconLink";
import Typed from "typed.js";

interface Props {}
interface State {
	mode: "light" | "dark";
}

export default class index extends Component<Props, State> {
	typed: Typed | undefined;

	constructor(props: Props) {
		super(props);
		this.state = {
			mode: "light",
		};
	}

	componentDidMount() {
		this.typed = new Typed("#type-b", {
			strings: ["better", "fast^300er", "smart^300er"],
			typeSpeed: 50,
			backSpeed: 50,
			loop: true,
			cursorChar: "&#9647;",
		});
		document.querySelector("body")?.classList.add(this.state.mode);
	}

	lightSwitch = () => {
		if (this.state.mode === "light") {
			this.setState({ mode: "dark" });
			document.querySelector("body")?.classList.add("dark");
			document.querySelector("body")?.classList.remove("light");
		} else {
			this.setState({ mode: "light" });
			document.querySelector("body")?.classList.add("light");
			document.querySelector("body")?.classList.remove("dark");
		}
	}

	render() {
		return (
			<>
				<Head>
					<title>Noteer, notes but better</title>
				</Head>
				<div className={`${styles["container"]}`}>
					<div className={styles["header"]}>
						<svg width="260" height="63" viewBox="0 0 260 63" fill="none" className={styles.logo} xmlns="http://www.w3.org/2000/svg">
							<path d="M70 14.8493H76.7553V42.2163H70V14.8493ZM79.435 31.8832C79.435 25.1148 83.5782 21.0763 89.9733 21.0763C96.3683 21.0763 100.512 25.1148 100.512 31.8832C100.512 38.6516 96.4584 42.7127 89.9733 42.7127C83.4882 42.7127 79.435 38.7644 79.435 31.8832ZM93.7337 31.8832C93.7337 28.1606 92.2701 25.9721 89.9733 25.9721C87.6765 25.9721 86.2353 28.2283 86.2353 31.8832C86.2353 35.5381 87.6539 37.7492 89.9733 37.7492C92.2926 37.7492 93.7337 35.6735 93.7337 31.9058V31.8832ZM103.056 43.5474H109.541C109.85 44.2344 110.372 44.8032 111.029 45.1692C111.686 45.5349 112.444 45.6784 113.189 45.578C115.756 45.578 117.107 44.1792 117.107 42.1486V38.3583H116.972C116.434 39.52 115.56 40.4933 114.465 41.153C113.369 41.8125 112.101 42.1279 110.825 42.0584C105.893 42.0584 102.628 38.2906 102.628 31.8155C102.628 25.3404 105.736 21.2568 110.915 21.2568C112.231 21.2051 113.532 21.5578 114.643 22.2673C115.753 22.9771 116.621 24.0099 117.13 25.2276V21.4598H123.885V41.9907C123.885 46.9317 119.539 50 113.122 50C107.154 50 103.439 47.3152 103.056 43.57V43.5474ZM117.13 31.8606C117.13 28.5667 115.621 26.491 113.234 26.491C110.847 26.491 109.384 28.5441 109.384 31.8606C109.384 35.1772 110.825 37.0498 113.234 37.0498C115.644 37.0498 117.13 35.2223 117.13 31.8832V31.8606ZM126.43 31.8606C126.43 25.0922 130.573 21.0537 136.968 21.0537C143.363 21.0537 147.529 25.0922 147.529 31.8606C147.529 38.6291 143.476 42.6901 136.968 42.6901C130.46 42.6901 126.43 38.7644 126.43 31.8832V31.8606ZM140.728 31.8606C140.728 28.138 139.265 25.9496 136.968 25.9496C134.671 25.9496 133.343 28.2283 133.343 31.9058C133.343 35.5833 134.761 37.7717 137.058 37.7717C139.355 37.7717 140.728 35.6735 140.728 31.9058V31.8606ZM150.096 16.5414C150.078 15.8581 150.263 15.1848 150.628 14.6071C150.993 14.0293 151.52 13.573 152.144 13.296C152.768 13.0191 153.459 12.9339 154.132 13.0513C154.804 13.1688 155.426 13.4835 155.919 13.9557C156.413 14.4278 156.755 15.0362 156.903 15.7035C157.051 16.3709 156.998 17.0671 156.75 17.7043C156.503 18.3412 156.072 18.8903 155.513 19.2818C154.953 19.6732 154.291 19.8896 153.609 19.9031C153.156 19.9342 152.703 19.8713 152.276 19.7183C151.849 19.5654 151.459 19.3258 151.129 19.0142C150.799 18.7028 150.537 18.3263 150.359 17.9084C150.181 17.4908 150.091 17.0407 150.096 16.5866V16.5414ZM150.096 21.505H156.851V42.2163H150.096V21.505ZM181.373 31.8606C181.373 38.6291 178.378 42.4871 173.244 42.4871C171.928 42.5717 170.617 42.2493 169.489 41.5634C168.361 40.8773 167.471 39.8609 166.939 38.6516H166.804V48.8494H160.049V21.4598H166.804V25.1599H166.939C167.447 23.9305 168.32 22.8877 169.44 22.1732C170.56 21.4587 171.872 21.1072 173.199 21.1665C178.378 21.2568 181.463 25.1373 181.463 31.9058L181.373 31.8606ZM174.618 31.8606C174.618 28.5667 173.109 26.4685 170.745 26.4685C168.38 26.4685 166.871 28.5892 166.849 31.8606C166.826 35.132 168.38 37.2302 170.745 37.2302C173.109 37.2302 174.618 35.1772 174.618 31.9058V31.8606ZM192.97 21.0312C198.577 21.0312 201.977 23.6934 202.134 27.9575H195.987C195.987 26.491 194.771 25.566 192.902 25.566C191.033 25.566 190.2 26.288 190.2 27.3484C190.2 28.4087 190.943 28.7472 192.452 29.063L196.775 29.9429C200.896 30.8228 202.652 32.4924 202.652 35.6284C202.652 39.9151 198.757 42.6675 193.015 42.6675C187.273 42.6675 183.512 39.9151 183.219 35.6961H189.727C189.93 37.2302 191.146 38.1327 193.127 38.1327C195.109 38.1327 196.009 37.4784 196.009 36.3955C196.009 35.3125 195.379 35.0869 193.758 34.7485L189.862 33.9137C185.831 33.079 183.715 30.9356 183.715 27.777C183.76 23.716 187.385 21.0763 192.97 21.0763V21.0312ZM225.193 42.1712H218.685V38.3132H218.55C218.192 39.6084 217.404 40.7422 216.314 41.5262C215.224 42.3099 213.9 42.6969 212.56 42.6224C211.534 42.6809 210.507 42.518 209.549 42.1446C208.591 41.7712 207.724 41.1963 207.006 40.4583C206.289 39.7206 205.738 38.8368 205.391 37.8674C205.044 36.8979 204.909 35.8651 204.994 34.8387V21.4598H211.749V33.282C211.749 35.7412 213.01 37.0498 215.105 37.0498C215.595 37.0507 216.082 36.9467 216.53 36.7443C216.976 36.5419 217.377 36.2459 217.701 35.8763C218.023 35.5066 218.266 35.0718 218.408 34.6012C218.55 34.1303 218.593 33.6344 218.527 33.1466V21.4598H225.283L225.193 42.1712ZM228.458 21.4598H234.988V25.4532H235.123C235.481 24.1751 236.253 23.052 237.316 22.2601C238.379 21.4682 239.676 21.0517 241 21.0763C242.349 20.9682 243.686 21.3723 244.751 22.2084C245.814 23.0445 246.526 24.2511 246.742 25.5886H246.877C247.285 24.2369 248.131 23.0601 249.282 22.2432C250.43 21.4262 251.82 21.0156 253.227 21.0763C254.146 21.0454 255.06 21.2085 255.911 21.555C256.765 21.9014 257.533 22.4237 258.17 23.0874C258.807 23.7514 259.298 24.5422 259.611 25.4083C259.922 26.2744 260.05 27.197 259.982 28.1154V42.2163H253.227V29.9204C253.227 27.6642 252.214 26.6264 250.322 26.6264C249.901 26.6228 249.482 26.7099 249.097 26.8818C248.712 27.0537 248.368 27.3064 248.089 27.6229C247.809 27.9395 247.6 28.3122 247.476 28.7163C247.352 29.1206 247.318 29.5465 247.373 29.9655V42.2163H240.955V29.8301C240.955 27.7545 239.942 26.6264 238.095 26.6264C237.67 26.6266 237.246 26.7173 236.859 26.8926C236.469 27.0681 236.12 27.3242 235.839 27.6439C235.555 27.9636 235.344 28.3399 235.215 28.7481C235.089 29.156 235.049 29.5862 235.1 30.0106V42.2163H228.345L228.458 21.4598Z" fill={this.state.mode !== "light" ? "#ebebeb" : "#394149"}/>
							<path d="M60.3991 19.0274C59.2173 16.3243 57.6532 13.8053 55.7546 11.5476C52.9481 8.11048 49.4457 5.30685 45.4779 3.32125C41.5101 1.33567 37.1668 0.21313 32.7341 0.0275065C28.3014 -0.158117 23.8797 0.597385 19.7598 2.24433C15.64 3.89125 11.9156 6.39225 8.83174 9.58272C5.74784 12.7732 3.37439 16.5808 1.8676 20.7549C0.360815 24.9289 -0.245175 29.375 0.0895539 33.8001C0.424306 38.2255 1.69214 42.5296 3.80965 46.4293C5.92717 50.3292 8.84635 53.7362 12.375 56.4262C14.6392 58.1867 17.1411 59.6178 19.8062 60.6771C23.5427 62.2088 27.5418 62.9977 31.58 63C39.9155 62.9816 47.9035 59.658 53.7934 53.7579C59.6831 47.8577 62.9939 39.8624 63 31.5246C63.006 27.2237 62.1205 22.9684 60.3991 19.0274ZM31.5568 6.41395C35.1832 6.41358 38.7657 7.20638 42.0533 8.73685C41.2517 9.08336 40.4087 9.32539 39.5453 9.45698C36.7043 9.86834 34.0725 11.1877 32.0426 13.2182C30.0127 15.2486 28.6937 17.8812 28.2824 20.7231C28.102 22.3168 27.3772 23.7995 26.2307 24.921C25.0842 26.0425 23.5861 26.7338 21.9891 26.8788C19.1481 27.2902 16.5162 28.6096 14.4864 30.64C12.4565 32.6705 11.1375 35.303 10.7263 38.1449C10.587 39.6604 9.94061 41.0845 8.89168 42.1868C7.08184 38.3616 6.26882 34.1404 6.52826 29.9165C6.7877 25.6925 8.11117 21.6028 10.3755 18.0278C12.6399 14.4531 15.7715 11.5093 19.479 9.47054C23.1865 7.4318 27.3493 6.36435 31.58 6.36749L31.5568 6.41395ZM12.3518 47.785C12.6305 47.5294 12.9092 47.2971 13.1878 47.0184C15.2538 45.0149 16.58 42.37 16.9499 39.5154C17.0998 37.9107 17.816 36.412 18.9702 35.2877C20.0809 34.1314 21.5758 33.4208 23.1735 33.29C26.0145 32.8786 28.6463 31.5592 30.6762 29.5288C32.706 27.4983 34.0251 24.8658 34.4364 22.0239C34.5369 20.3456 35.2568 18.7644 36.4567 17.5872C37.5948 16.4532 39.1038 15.7685 40.7064 15.6591C43.3164 15.2999 45.7506 14.1393 47.6732 12.3374C49.8623 14.1541 51.7169 16.3395 53.1537 18.7951C53.0654 18.9142 52.9642 19.0234 52.8518 19.1203C51.7434 20.2868 50.2504 21.013 48.6485 21.1644C45.8044 21.5686 43.1687 22.8857 41.1374 24.9173C39.1064 26.9492 37.7897 29.5857 37.3856 32.4305C37.2219 34.0296 36.5106 35.523 35.3722 36.6575C34.2336 37.7923 32.7381 38.4982 31.1388 38.6559C28.2894 39.0641 25.6478 40.3809 23.6061 42.4104C21.5646 44.4402 20.2319 47.0744 19.8062 49.922C19.6813 51.0572 19.2914 52.1474 18.6683 53.1044C16.2857 51.6893 14.1518 49.8923 12.3518 47.785ZM31.5568 56.7514C29.2004 56.7539 26.8552 56.4255 24.59 55.7758C25.365 54.3309 25.8687 52.7562 26.0763 51.1299C26.2391 49.5343 26.9471 48.0437 28.0808 46.9097C29.2146 45.7756 30.7047 45.0674 32.2999 44.9046C35.1451 44.4978 37.7822 43.1805 39.8168 41.1496C41.851 39.1189 43.1736 36.4835 43.586 33.6384C43.7488 32.0428 44.4568 30.5522 45.5906 29.4182C46.7243 28.2841 48.2145 27.5759 49.8096 27.4131C51.9507 27.1345 53.9915 26.3371 55.7546 25.0901C56.7336 28.8019 56.8495 32.6886 56.0936 36.4522C55.3377 40.2158 53.7298 43.7563 51.3936 46.8019C49.0574 49.8472 46.055 52.3172 42.6162 54.0217C39.1777 55.7263 35.3945 56.6204 31.5568 56.6352V56.7514Z" fill={this.state.mode !== "light" ? "#ebebeb" : "#394149"}/>
						</svg>
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
								{
									this.state.mode === "light" ? (
										<BsMoon className={styles["modeswitch"]} onClick={this.lightSwitch}></BsMoon>
									) : (
										<FiSun className={styles["modeswitch"]} onClick={this.lightSwitch}></FiSun>
									)
								}
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
									<p className={styles["title"]}>Notes, but </p>
									<div className={styles["code-title"]}><div id="type-b"></div></div>
								</div>
								<div className={styles["content-right-bottom"]}>
									<div className={styles["benefit-right"]}>
										<span className={styles["text-bold"]}>Better</span>
										<p className={styles["text-normal"]}>
											Take better notes, and take your memory and knowledge to
											another level. Decreasing the gap between ideas and words
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
											note-taker without making much effort
										</p>
									</div>
									<div className={styles["tryit"]}>
										<button className={styles["primary-btn"]}>
											Try it! It&lsquo;s free!
										</button>
										<div className={`${styles["tryit-sh"]} ${styles["text-normal"]}`}>
											(and self hosted if you wish)
										</div>
									</div>
								</div>
							</div>
						</div>
						<div className={styles["content-bottom"]}>
							<span className={styles["subtitle"]}>Make your notes useful</span>
						</div>
					</div>
				</div>
			</>
		);
	}
}

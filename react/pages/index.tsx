import type { NextPage } from "next";
import Head from "next/head";
import Image from "next/image";
import styles from "../styles/Home.module.css";

const Home: NextPage = () => {
  return (
    <div className={styles.container}>
      <Head>
        <title>Seclusion</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <div>
        <h1>Seclusion</h1>
        <p>Seclusion is a place to share your thoughts and ideas.</p>
        <a href="/login">Login</a>
      </div>
    </div>
  );
};

export default Home;

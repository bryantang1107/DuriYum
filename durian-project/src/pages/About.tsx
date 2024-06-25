import { useEffect } from "react";
import axios from "../axios";
const About = () => {
  useEffect(() => {
    const getData = async () => {
      try {
        const response = await axios.get("/ping", {
          headers: {
            Authorization: `Basic ${btoa("user:password")}`,
          },
        });
        console.log(response);
      } catch (error) {
        console.log(error);
      }
    };

    getData();
  }, []);
  return <div>About</div>;
};

export default About;

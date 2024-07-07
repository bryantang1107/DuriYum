import Container from "../../components/Container";
import Col from "../../components/Col";
import { useState } from "react";
import UserProfile from "./UserProfile";
import Order from "./Order";
import Button from "../../components/Button";
import Main from "../../components/Main";

const Profile = () => {
  const [showProfile, setShowProfile] = useState<boolean>(true);

  return (
    <Main>
      <Container isCenter={true}>
        <Container className="bg-gray-50 flex-col max-w-max p-10">
          <h1 className="text-center font-semibold">Your Profile</h1>
          <Container className="w-full justify-around mt-10 gap-10">
            <Col size={3}>
              <Container className="flex-col gap-3">
                <Button
                  onClick={() => setShowProfile(true)}
                  label="MY ACCOUNT"
                  className={`${showProfile && "font-semibold"}`}
                ></Button>
                <Button
                  onClick={() => setShowProfile(false)}
                  label="ORDERS"
                  className={`${!showProfile && "font-semibold"}`}
                ></Button>
              </Container>
            </Col>
            <Col size={5}>
              {showProfile ? <UserProfile></UserProfile> : <Order></Order>}
            </Col>
          </Container>
        </Container>
      </Container>
    </Main>
  );
};

export default Profile;

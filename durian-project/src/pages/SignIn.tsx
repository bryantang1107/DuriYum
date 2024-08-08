import { useRef } from "react";
import Form from "../components/Form";
import Input from "../components/Input";
import Label from "../components/Label";
import Button from "../components/Button";
import Container from "../components/Container";
import Row from "../components/Row";
import Main from "../components/Main";
import { useAuth } from "../context/AuthContext";

const SignIn = () => {
  const emailRef = useRef<HTMLInputElement>(null);
  const passwordRef = useRef<HTMLInputElement>(null);
  const {firebaseUtil} = useAuth();
  const handleSignIn =  async () => {
    //normal login
    if(emailRef.current && passwordRef.current){
      await firebaseUtil.signIn(emailRef.current.value, passwordRef.current.value)
    }
  };

  const loginWithGoogle =  async () =>{
    //login with google
    await firebaseUtil.signInWithGoogle()
  }
  return (
    <Main>
      <Container isCenter={true}>
        <div className="flex flex-col w-3/5">
          <h1 className="text-center text-3xl mb-8">Sign In </h1>
          <p className="italic mb-5 text-center">
            Lorem ipsum dolor sit amet, consectetur adipscing
          </p>
          <Form>
            <Row className="flex flex-col">
              <Label label="EMAIL" className="font-semibold"></Label>
              <Input
                ref={emailRef}
                placeholder="Enter Email ..."
                className="bg-gray-50 border border-sky-500"
                type="text"
              ></Input>
            </Row>
            <Row className="flex flex-col">
              <Label label="PASSWORD" className="font-semibold"></Label>
              <Input
                ref={passwordRef}
                placeholder="Enter Password ..."
                className="bg-gray-50 border border-sky-500"
                type="text"
              ></Input>
            </Row>
            <Container className="items-center gap-2">
              <Input type="checkbox"></Input>
              <h3>REMEMBER ME</h3>
            </Container>
            <Button
              label="SIGN IN"
              onClick={handleSignIn}
              className="rounded-full bg-gray-50 border border-red-200 p-2 mx-auto block mt-2"
            ></Button>
          </Form>
          <Button label="LOGIN WITH GOOGLE" onClick={loginWithGoogle}>
          </Button>
          <Container className="flex-col items-center">
            <p>
              Don't have an account ?
              <a href="/" className="underline">
                Sign Up
              </a>
            </p>
            <p>
              <a href="/" className="underline">
                Forgot your password ?
              </a>
            </p>
          </Container>
        </div>
      </Container>
    </Main>
  );
};

export default SignIn;

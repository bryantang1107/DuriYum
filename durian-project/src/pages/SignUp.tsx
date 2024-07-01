import { useRef } from "react";
import Form from "../components/Form";
import Input from "../components/Input";
import Label from "../components/Label";
import Button from "../components/Button";
import Container from "../components/Container";
const SignUp = () => {
  const emailRef = useRef<HTMLInputElement>(null);
  const passwordRef = useRef<HTMLInputElement>(null);
  const handleSignUp = () => {};
  return (
    <div className="flex items-center justify-center h-screen">
      <div>
        <h1 className="text-center text-3xl mb-8">Sign Up </h1>
        <p className="italic mb-5 text-center">
          Lorem ipsum dolor sit amet, consectetur adipscing
        </p>
        <Form>
          <Container className="flex-col">
            <Container className="gap-2">
              <Container>
                <Label label="FIRST NAME" className="font-semibold"></Label>
                <Input
                  className="mx-2 bg-gray-50 border border-sky-500"
                  type="text"
                  placeholder="e.g. John Doe"
                ></Input>
              </Container>
              <Container>
                <Label label="LAST NAME" className="font-semibold"></Label>
                <Input
                  className="mx-2 bg-gray-50 border border-sky-500"
                  type="text"
                  placeholder="e.g. John Doe"
                ></Input>
              </Container>
            </Container>
            <Label label="EMAIL" className="font-semibold"></Label>
            <Input
              ref={emailRef}
              placeholder="Enter Email ..."
              className="bg-gray-50 border border-sky-500"
              type="text"
            ></Input>
            <Label label="PASSWORD" className="font-semibold"></Label>
            <Input
              ref={passwordRef}
              placeholder="Enter Password ..."
              className="bg-gray-50 border border-sky-500"
              type="text"
            ></Input>
          </Container>
          <Button
            label="CREATE ACCOUNT"
            onClick={handleSignUp}
            className="rounded-full bg-gray-50 border border-red-200 p-2 mx-auto block mt-4"
          ></Button>
          <Container className="flex-col items-center">
            <p>
              Have an account ?
              <a href="/" className="underline">
                Login
              </a>
            </p>
          </Container>
        </Form>
      </div>
    </div>
  );
};

export default SignUp;

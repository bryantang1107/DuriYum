import Container from "../../components/Container";
import Form from "../../components/Form";
import Input from "../../components/Input";
import Label from "../../components/Label";
import Row from "../../components/Row";

const UserProfile = () => {
  return (
    <Form>
      <Container className="flex-col gap-5">
        <h1 className="font-semibold">PERSONAL DETAILS</h1>
        <Row className="flex-col">
          <Label label="NAME" className="font-semibold"></Label>
          <Input type="text" placeholder="John Doe"></Input>
        </Row>
        <Row className="flex-col">
          <Label label="EMAIL" className="font-semibold"></Label>
          <Input type="text" placeholder="johndoe@email.com"></Input>
        </Row>
        <Row className="flex-col">
          <Label label="PASSWORD" className="font-semibold"></Label>
          <Input type="text" placeholder="********"></Input>
        </Row>
      </Container>
      <Container className="mt-36 flex-col">
        <h1 className="font-semibold">ADDRESSES</h1>
        <Row className="mt-5">
          <Label label="DEFAULT ADDRESS" className="font-semibold"></Label>
          <Input
            type="textarea"
            placeholder="No.1 Jalan SS2/3..."
            row={4}
            col={50}
          ></Input>
        </Row>
        Icon to add
      </Container>
    </Form>
  );
};

export default UserProfile;

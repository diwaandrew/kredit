import React, { Component } from 'react'
import { Image,Col,Row,Container} from "react-bootstrap";
import BankSinarmas from '../../assets/BankSinarmas.png'
import './Home.css'
export default class Home extends Component {
  render() {
    return (
        <Container fluid className=''>
            <Row className='d-flex justify-content-center align-items-center  '>
                <Col col='12' >
                    <div className='mx-auto content p-5 text-center home ' >
                    <hr></hr>
                    <h1>Welcome Program Bantu Kredit</h1>
                    <h4>Supported By :</h4>
                      <Image src={BankSinarmas} width="300" />
                    <hr></hr>
                    </div>
                </Col>
            </Row>
        </Container>
    )
  }
}
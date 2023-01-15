import React, { Component } from 'react'
import { Button,Image,Col,Row,Container,Card,Form, InputGroup } from "react-bootstrap";
import BankSinarmas from '../../assets/BankSinarmas.png'
import './Login.css'
import * as RiIcons from "react-icons/ri";
import swal from "sweetalert";

export default class Login extends Component {
    handleSubmit = async(event) => {
        localStorage.setItem("name","john"); 
        if (localStorage.getItem("name")==="john"){
            swal({
                title: "Sukses Login",
                text: "Welcome " +localStorage.getItem("name") ,
                icon: "success",
                button : false,
                timer : 1500,
            }).then(()=>{ 
                window.location.href="/"   
            })
        }else{
            swal({
                title: "Gagal Login",
                text: "User Login Tidak Sesuai" ,
                icon: "error",
                button : false,
                timer : 1500,
            })
        }
        
	};
    render() {
        return (
            <div className='login-body'>
                <Container fluid className='login-container mb-3'>
                    <Row className='d-flex justify-content-center align-items-center'>
                        <Col col='12' >
                            <div className='mx-auto log-shadow p-5' >
                                <Image src={BankSinarmas} width="300" height="200" className='text-center'/>
                                <Card.Body className='w-100 d-flex flex-column'>
                                    <Row>
                                    <InputGroup  className="btn-shadow mb-2" >
                                        <InputGroup.Text id="basic-addon1" className='btn-input'>
                                            <RiIcons.RiUser3Line />
                                        </InputGroup.Text>
                                        <Form.Control
                                            className='btn-input'
                                            placeholder="Username"
                                            aria-label="Username"
                                            aria-describedby="basic-addon1"
                                        />
                                    </InputGroup>
                                    <InputGroup className="btn-shadow mb-3" >
                                        <InputGroup.Text className='btn-input' id="basic-addon1">
                                            <RiIcons.RiLockPasswordLine />
                                        </InputGroup.Text>
                                        <Form.Control
                                            className='btn-input'
                                            placeholder="Password"
                                            aria-label="Password"
                                            aria-describedby="basic-addon1"
                                            type ="password"
                                        />
                                    </InputGroup>
                                    <Button className='btn-shadow btn-login' variant='Primary' style={{ backgroundColor:"#128297",color:"white"}} onClick={(e)=>this.handleSubmit(e)}>
                                            Login
                                    </Button>    
                                    </Row>
                                </Card.Body>
                            </div>
                        </Col>
                    </Row>
                </Container>
            </div>                
        )
    }
}
import React, { Component } from 'react'
import { Button,Col,Row,Container,Card,Form, InputGroup } from "react-bootstrap";
import { API_URL } from '../../const'
import './ChangePassword.css'
import axios from 'axios'
import swal from 'sweetalert'
import * as RiIcons from "react-icons/ri";

export default class ChangePassword extends Component {
    constructor(props) {
		super(props);
		this.state ={
            oldpassword : "",
            password : "", 
            confirmpassword:"", 
            change:[],
        }  
	}
    handleoldpassword = e => {
        const { value } = e.target;
        this.setState({oldpassword : value});
    };    
    handlePassword = e => {
        const { value } = e.target;
        this.setState({ password : value});
    };
    handleConfirmPassword = e => {
        const { value } = e.target;
        this.setState({ confirmpassword : value});
    };
    handleSubmit = async(event) => {

        const change = 
        {
            nik:localStorage.getItem("nik"),
            oldpassword:    this.state.oldpassword,
            password:this.state.password
        }
        if(this.state.password !== this.state.confirmpassword){
            swal({
                title: "Oops Something went wrong",
                text: "New Password and Confirm New Password are not match" ,
                icon: "error",
                button : false,
                timer : 2000,
        })
        }else{
            axios
			.put(API_URL+"/updatePassword",change)
			.then(res => {
                swal({
                    title: "Success",
                    text: "Change Password Success" ,
                    icon: "success",
                    button : false,
                    timer : 1000,
            }).then(()=>{ 
                window.location.href="/changepassword"   
            })
            }).catch(error => 
                swal({
                    title: "Error",
                    text: "Old password is wrong" ,
                    icon: "error",
                    button : false,
                    timer : 1000,
            }));  
        }
    };
    render() {
        return (
            <Container fluid className=''> 
                <Row className='d-flex justify-content-center align-items-center'>
                    <Col col='12' >
                        <div className='mx-auto content p-5' >
                            <div className='login-body'>
                                <Container fluid className='login-container mb-3'>
                                    <h2 className='d-flex justify-content-center align-items-center pb-4'>Change Password</h2>
                                    <Row className='d-flex justify-content-center align-items-center'>
                                        <Col col='12' >
                                            <div className='mx-auto box p-5' >
                                                <Card.Body className='w-100 d-flex flex-column mt-4'>
                                                    <Row>
                                                    <InputGroup  className="btn-shadow mb-3" onChange={this.handleoldpassword} value={this.state.oldpassword}>
                                                        <InputGroup.Text id="basic-addon1" className='btn-input'><RiIcons.RiLockPasswordLine /></InputGroup.Text>
                                                        <Form.Control
                                                        className='btn-input'
                                                        placeholder="Old Password"
                                                        aria-label="Old Password"
                                                        aria-describedby="basic-addon1"
                                                        type ="password"
                                                        />
                                                    </InputGroup>
                                                    <InputGroup className="btn-shadow mb-3" onChange={this.handlePassword} value={this.state.password}>
                                                        <InputGroup.Text className='btn-input' id="basic-addon1"><RiIcons.RiLockUnlockLine /></InputGroup.Text>
                                                        <Form.Control
                                                        className='btn-input'
                                                        placeholder="New Password"
                                                        aria-label="New Password"
                                                        aria-describedby="basic-addon1"
                                                        type ="password"
                                                        />
                                                    </InputGroup>
                                                    <InputGroup className="btn-shadow mb-3" onChange={this.handleConfirmPassword} value={this.state.confirmpassword} >
                                                        <InputGroup.Text className='btn-input' id="basic-addon1"><RiIcons.RiLockLine /></InputGroup.Text>
                                                        <Form.Control
                                                        className='btn-input'
                                                        placeholder="Confirm New Password"
                                                        aria-label="Confirm New Password"
                                                        aria-describedby="basic-addon1"
                                                        type ="password"
                                                        />
                                                    </InputGroup>
                                                        <Row>
                                                            <Button className='btn-shadow btn-login' variant='Primary' style={{ backgroundColor:"#128297",color:"white"}} onClick={(e)=>this.handleSubmit(e)}>
                                                                Submit
                                                            </Button> 
                                                        </Row>   
                                                    </Row>
                                                </Card.Body>
                                            </div>
                                        </Col>
                                    </Row>
                                </Container>
                            </div>
                        </div>
                    </Col>
                </Row>
            </Container>
        )
    }
}
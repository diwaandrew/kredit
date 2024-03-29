import React, { Component } from 'react'
import DataTable from 'react-data-table-component'
import axios from 'axios'
import swal from 'sweetalert'
import { API_URL } from '../../const'
import { 
    Button, 
    FormGroup, 
    FormSelect, 
    FormControl,
    Form, 
    Container,
    Col,
    Row
} from 'react-bootstrap';
import * as FaIcons from "react-icons/fa";
import Moment from 'moment';
import { FormatRupiah } from "@arismun/format-rupiah";


export default class Checklist extends Component {
    constructor(props){
        super(props)
        this.state = { customers:[], checked:[], branch:[], company:[],data:[],
            currentDate:new Date().toISOString().split('T')[0], isSubmit:false
        };
    }

    componentDidMount(){
        axios
            .get(API_URL+"/listreport?statustrx=0")
            .then(res => {
                const customers = res.data.data;
                this.setState({customers});
                console.log(customers)
            })
            .catch(error => console.log(error));
        
        axios
            .get(API_URL+"/branch")
            .then(res => {
                const branch = res.data.data;
                this.setState({branch});
                // console.log(branch)
            })
            .catch(error => console.log(error));
        
        axios
            .get(API_URL+"/company")
            .then(res => {
                const company = res.data.data;
                this.setState({company});
                // console.log(company)
            })
            .catch(error => console.log(error));
    }

    checklist = (ppk,event) => {
        if(event.target.checked){
            const checked = [...this.state.checked,{Ppk:ppk}]
            this.setState({checked})
            // console.log(checked)
        }else{          
            let checkedData = this.state.checked
            checkedData=checkedData.filter((j)=> j.Ppk !== ppk)
            this.setState({checked:checkedData})
            // console.log(checkedData)
        }
    }

    updateApproval = () => {
        // console.log(this.state.checked)
        if(this.state.checked.length === 0){
            swal({
                title: "Oops Something went wrong   ",
                text: "Choose Data First !" ,
                icon: "info",
                button : false,
                timer : 1500,
            })
        }
        else{
            axios
            .put(API_URL+"/updateflag", this.state.checked)
            swal({
                title: "Approve Sukses",
                text: "Approve" ,
                icon: "success",
                button : false,
                timer : 1500,
            }).then(()=>{ 
            window.location.href="/approval"   
            })
        }
    }

    handleSubmit = async(event) => {
        event.preventDefault();
        const formData = new FormData(event.currentTarget);
        // console.log(formData.get('branch'))
        if(formData.get('branch')==="Please Choose"){
            swal({
                    title: "Oops Something went wrong   ",
                    text: "Choose Branch First !" ,
                    icon: "info",
                    button : false,
                    timer : 1000,
            })
        }else if(formData.get('company')==="Please Choose"){
            swal({
                title: "Oops Something went wrong   ",
                text: "Choose Company First !" ,
                icon: "info",
                button : false,
                timer : 1000,
            })
        }else{
            this.setState({isSubmit:true})
            axios
                .get(API_URL+"/search?branch="+formData.get('branch')+"&company="+formData.get('company')+"&startdate="+formData.get('startDate')+"&enddate="+formData.get('endDate')+"&statustrx=9")
                .then((res) => {
                    const customers = res.data.data;
                    this.setState({ customers });
                    // console.log(customers);
                })
                .catch((error) => {
                    console.log("Error yaa ", error);
                });
        }
    }
    render() {
        let branchList = this.state.branch.map(
            (branchList)=>(
                <option value={branchList.code}>
                    {branchList.code}&nbsp;&nbsp;-&nbsp;&nbsp;{branchList.description}
                </option>
            )
        )

        let companyList = this.state.company.map(
            (companyList)=>(
                <option value={companyList.company_short_name}>
                    {companyList.company_short_name}
                </option>
            )
        )
        let data = []
        this.state.customers.map(
            (customerList,id)=>(
                data = [...data,{
                    id:id+1, 
                    Ppk:customerList.Ppk, 
                    Name:customerList.Name,
                    Company:customerList.Company,
                    DrawdownDate:Moment(customerList.DrawdownDate).format('LLL'),
                    Loan_Amount:<FormatRupiah value={customerList.Loan_Amount}/>,
                    InterestEffective:customerList.InterestEffective
                }]
            )
        )
        const kolom = [
            {
                name : "No",
                selector: row => row.id, 
                sortable: true, 
                width: '70px', 
            },
            {
                name : "PPK",
                selector: row => row.Ppk,   
            },
            {
                name : "Name",
                selector: row => row.Name, 
                sortable: true,  
            },
            {
                name : "Company",
                selector: row => row.Company,   
            },
            {
                name : "Drawdown Date",
                selector: row => row.DrawdownDate,  
            },
            {
                name : "Loan Amount",
                selector: row => row.Loan_Amount,   
                sortable: true,
            },
            {
                name : "Interest Effective",
                selector: row => row.InterestEffective,
                sortable: true,   
            },
        ]
    
        return (
            <Container fluid>
                <Form onSubmit={(e)=>this.handleSubmit(e)}>  
                    <Row ><h2 className='d-flex justify-content-center align-items-center pb-4'>Drawdown Report</h2> </Row>
                    <Row className='d-flex justify-content-center align-items-center pb-4'>
                        <Col xs="auto">
                            <Row xs="auto">
                                <Col className='pt-2'>
                                    <label>Branch :</label>
                                </Col>
                                <Col>
                                    <FormGroup>  
                                        <FormSelect name='branch'>
                                            <option className='d-none'>Please Choose</option>
                                            <option value='' >All Branch</option>
                                            {branchList}
                                        </FormSelect>
                                     </FormGroup>
                                </Col>
                            </Row>
                        </Col>
                        <Col xs="auto">
                            <Row xs="auto">
                                <Col className='pt-2'>
                                    <label>Company :</label>
                                </Col>
                                <Col>
                                    <FormGroup>  
                                        <FormSelect name='company'>
                                            <option className='d-none'>Please Choose</option>
                                            <option value=''>All Company</option>
                                        {companyList}
                                        </FormSelect>
                                    </FormGroup>
                                </Col>
                            </Row>
                        </Col>
                        <Col xs="auto">
                            <Row xs="auto">
                                <Col className='pt-2'>
                                    <label>Start :</label>
                                </Col>
                                <Col>
                                    <FormGroup>
                                        <FormControl type='date' name='startDate' defaultValue={this.state.currentDate}></FormControl>
                                    </FormGroup>
                                </Col>
                            </Row>
                        </Col>
                        <Col xs="auto">
                            <Row xs="auto">
                                <Col className='pt-2'>
                                    <label>End :</label>
                                </Col>
                                <Col>
                                    <FormGroup>
                                        <FormControl type='date' name='endDate' defaultValue={this.state.currentDate}></FormControl>
                                    </FormGroup>
                                </Col> 
                            </Row>
                        </Col>
                        <Col xs="auto">
                            <Button type='submit' variant="outline-info">
                                <FaIcons.FaSearch/> Submit
                            </Button>
                        </Col>
                    </Row>
                    <Row></Row>
                    <Row>
                        <DataTable 
                            // title='Approval'
                            columns={kolom}
                            data={data}
                            direction="auto"
                            fixedHeaderScrollHeight="300px"
                            pagination
                            responsive
                            subHeaderAlign="right"
                            subHeaderWrap
                        />
                    </Row>
                    <Row>
                        {/* <Col xs="auto">
                            <Button onClick={() => this.updateApproval()} hidden={this.state.customers.length === 0 || this.state.checked.length  === 0} >
                               <FaIcons.FaCheck /> Approve
                            </Button>
                        </Col> */}
                    </Row>
                </Form>
            </Container>    
        )
    }
}

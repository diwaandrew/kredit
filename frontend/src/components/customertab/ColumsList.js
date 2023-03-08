export const ColumnsList = [
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
    {
        name : "Action",
        selector: row => row.action,
        center : true,   
    },
]
import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
 
import moment from 'moment';

const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
import { EntBill } from '../../api/models/EntBill';


export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 const [bills, setBills] = useState<EntBill[]>([]);
 const [loading, setLoading] = useState(true);
 
 useEffect(() => {
   
  const getBills = async () => {
    const res = await api.listBill({ limit: 10, offset: 0 });
    setLoading(false);
    setBills (res);
    console.log(res);
  };
  getBills ();
 }, [loading]);
 
 const deleteBills = async (id: number) => {
   const res = await api.deleteBill({ id: id });
   setLoading(true);
 };
 
 return (
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">No.</TableCell>
           <TableCell align="center">Product</TableCell>
           <TableCell align="center">Quantity</TableCell>
           <TableCell align="center">Seller</TableCell>
           <TableCell align="center">Customer</TableCell>
           <TableCell align="center">Time</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         {bills.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges.product.productName}</TableCell>
             <TableCell align="center">{item.quantity}</TableCell>
             <TableCell align="center">{item.edges.user.userEmail}</TableCell>
             <TableCell align="center">{item.edges.customer.customerName}</TableCell>
             <TableCell align="center">{moment(item.addedTime).format('DD/MM/YYYY HH:mm')}</TableCell>
             <TableCell align="center">
               <Button
                 onClick={() => {
                   deleteBills(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
               >
                 Delete
               </Button>
             </TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
 );
}

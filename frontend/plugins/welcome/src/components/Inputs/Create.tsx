import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';

import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';

import { EntCustomer } from '../../api/models/EntCustomer';
import { EntProduct } from '../../api/models/EntProduct';
import { EntUser } from '../../api/models/EntUser';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(3),
    },
    withoutLabel: {
      marginTop: theme.spacing(3),
    },
    textField: {
      width: '25ch',
    },
  }),
);


export default function Create() {
  const classes = useStyles();
  const profile = { givenName: 'กรอกข้อมูลรายการสินค้า' };
  //const username = { givenuser: 'b6108526@g.sut.ac.th' };
  //const logout = { givenlogout: 'logout' };
  const api = new DefaultApi();

  const [products, setProducts] = useState<EntProduct[]>([]);
  const [users, setUsers] = useState<EntUser[]>([]);
  const [customers, setCustomers] = useState<EntCustomer[]>([]);
 
  //const [bills, setBills] = useState<EntBill[]>([]);
  


  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);

  const [loading, setLoading] = useState(true);
  const [time, setTime] = useState(String);

  const [productid, setProduct] = useState(Number);
  
  const [userid, setUser] = useState(Number);
  const [customerid, setCustomer] = useState(Number);

  const [quantityid, setQuantity] = useState(String);
  let quantity = Number(quantityid);
  console.log(quantity);

  //const [billid, setBill] = useState(Number);

  useEffect(() => {
    const getProducts = async () => {
      const res = await api.listProduct({ limit: 10, offset: 0 });
      setLoading(false);
      setProducts(res);
    };
    getProducts();

    const getUsers = async () => {
      const res = await api.listUser({ limit: 10, offset: 0 });
      setLoading(false);
      setUsers(res);
      console.log(res);
    };
    getUsers();

    const getCustomers = async () => {
      const res = await api.listCustomer({ limit: 10, offset: 0 });
      setLoading(false);
      setCustomers(res);
    };
    getCustomers();

  }, [loading]);


    const handleTimeChange = (event: any) => {
      setTime(event.target.value as string);
    };
    const handleQuantityChange = (event: any) => {
      setQuantity(event.target.value);
    };
    const ProducthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setProduct(event.target.value as number);
    };

    const UserhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setUser(event.target.value as number);
    };

    const CustomerhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setCustomer(event.target.value as number);
    };

  

  const CreateBill = async () => {
    const bill = {
      productID: productid,
      userID: userid,
      customerID: customerid,
      quantityID: quantity,
      addtime: time + ":00+07:00",
  
    };
    console.log(bill);
    const res = await api.createBill({ bill:bill });
    setStatus(true);
    if (res.id != '') {
      setAlert(true);
    } else {
      setAlert(false);
    }
    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);
  };

  return (
    <Page theme={pageTheme.home}>
      <Header
        title={` ${profile.givenName || 'to Backstage'}`}
        subtitle=""
        
      ></Header>
      <Content>
      <form className={classes.margin} noValidate align ="right">
           <Button variant="contained" component={RouterLink} to="/" >
             Log Out
           </Button>
        </form>
        <ContentHeader title="">
          
          {status ? (
            <div>
              {alert ? (
                <Alert severity="success">
                  This is a success alert — check it out!
                </Alert>
              ) : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    This is a warning alert — check it out!
                  </Alert>
                )}
            </div>
          ) : null}
        </ContentHeader>

        <div className={classes.root}>
          <form noValidate autoComplete="off" align ="center">

            <div>
              <FormControl
                //fullWidth
                className={classes.margin}
                variant="outlined"
                style={{ width: 500 }}
              >
                 <InputLabel id="product-label">Product</InputLabel>
                <Select

                  label="Product"
                  variant="outlined"
                  value={productid}
                  onChange={ProducthandleChange}
                  
                >
                {products.map((item: EntProduct) => (
                  <MenuItem value={item.id}>{item.productName}</MenuItem>
                ))}
                </Select>

              </FormControl>
            </div>
            
            <div>
              <FormControl
                //fullWidth
                className={classes.margin}
                variant="outlined"
                style={{ width: 500 }}
              >
                <TextField
                  id="quantity"
                  label="Quantity"
                  variant="outlined"
                  type="number"
                  size="medium"
                  value={quantityid}
                  onChange={handleQuantityChange}
                />
              </FormControl>
            </div>

            <div>
              <FormControl
                //fullWidth
                className={classes.margin}
                variant="outlined"
                style={{ width: 500 }}
              >
                <InputLabel id="user-label">Seller</InputLabel>
                <Select
                labelId="user-label"
                label="Seller"
                variant="outlined"
                id="seller"
                value={userid}
                onChange={UserhandleChange}
                
              >
                {users.map((item: EntUser) => (
                  <MenuItem value={item.id}>{item.userEmail}</MenuItem>
                ))}
              </Select>

              </FormControl>
            </div>

            <div>
              <FormControl
                //fullWidth
                className={classes.margin}
                variant="outlined"
                style={{ width: 500 }}
              >
                <InputLabel id="customer-label">Customer</InputLabel>
                <Select
                labelId="customer-label"
                label="Customer"
                variant="outlined"
                id="customer"
                value={customerid}
                onChange={CustomerhandleChange}
                
              >
                {customers.map((item: EntCustomer) => (
                  <MenuItem value={item.id}>{item.customerName}</MenuItem>
                ))}
              </Select>

              </FormControl>
            </div>

            <form className={classes.margin} noValidate>
              <TextField
                id="datetime-local"
                label="Next appointment"
                type="datetime-local"
                value={time}
                onChange={handleTimeChange}
                className={classes.textField}
                InputLabelProps={{
                  shrink: true,
                }}
              />
            </form>
            <div className={classes.margin}>
              <Button
                onClick={() => {
                  CreateBill();
                }}
                variant="contained"
                color="primary"
              >
                Submit
             </Button>
              <Button
                style={{ marginLeft: 20 }}
                component={RouterLink}
                to="/output"
                variant="contained"
              >
                Back
             </Button>
            </div>
          </form>
        </div>
      </Content>
    </Page>
  );
}
Need to install `abigen` then create a
Go package from the `StaticOrganization.sol`.

> Installation guide:
> https://geth.ethereum.org/docs/tools/abigen
> 
> Once the abi is installed, add the directory with abigen
> to the environment path.

In the root directory of the proxy organization run the following
after each re-deployment of the smartcontract.
```bash
abigen  
  --abi=D:/sds/static-organization/abi/StaticOrganization.json 
  --pkg=smartcontract 
  --out=smartcontract/static_organization.go 
  --type StaticOrganization
```
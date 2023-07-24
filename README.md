# Organization proxy
Organization Proxy does solve two tasks:
* Convert the mysql commands to the w3storage commands.
* Lints the [IPFS](https://ipfs.io/) content ids to the topic in the smartcontract.

Set the **organization proxy** between **static service** and **w3storage extension**.

What does **organization proxy** solve?

**static service** doesn't know the storage internal structures.
The storages are bind through *extension*.
For **static service** they are all the same:
* whether it's [mysql](https://mysql.com).
* or it's [ipfs](https://IPFS.io)

However, storages are using different mechanism to store the content.
**organization proxy** creates *CID* and *file name* from the content parameters.
The **w3storage extension** identifies the *CID* and *file name*.
But **static service** doesn't know *file name* or *CID*.

For the developers to find the files on [IPFS](https://ipfs.io/) the *CID* are set on smartcontract.

## Dependency

Requires the smartcontract [StaticOrganization](https://github.com/ahmetson/static-organization).

Need to install `abigen`.
> Installation guide:
> https://geth.ethereum.org/docs/tools/abigen
> 
> Once the abi is installed, add the directory with abigen
> to the environment path.

For working with the smartcontract, let's create a .go.

Go to the root directory of the proxy organization.
Run the following.
```bash
abigen  
  --abi=D:/sds/static-organization/abi/StaticOrganization.json 
  --pkg=smartcontract 
  --out=smartcontract/static_organization.go 
  --type StaticOrganization
```

It will require the smartcontract to be installed on your computer.

> :warning:
> Re-generate the go abi after each redeployment of the contract.
# Branches CLI

## Entities

### Command
**Property**
- Name `string`

### Region
**Property** 
- Id `int`
- Name `string`
- Branches `[]int`

**Methods**
- List() `[]Branches`
- CountOfBranches() int
- CountOfEmployees() int

### Branch
**Property**
- Id `int`
- Name `string`
- Phone `string`
- CreateDate `DateTime`
- NumberOfEmployees `int`

**Methods**
- Get(int id) Branch
- New(name string, phone string, date string, numOfEmp int) Branch
- Edit(name string, phone string, date string, numOfEmp int) Branch



## User Story

- `list` `Branches` in a region
- `get` a `Branch` detail with `Id` 
- `create` a `Branch`
- `edit` a `Branch`
- a `status` for `Region` and `count` of `branches` and count of `employees`
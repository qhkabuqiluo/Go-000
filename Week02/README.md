### 作业
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么？应该怎么做请写出代码
### 个人理解
Dao层不应该Wrap这个error，而应该Service层Wrap，sql.ErrNoRows从某种角度对于Dao层来说应该是一个正常的返回，由于业务层的需要，应该在业务层去Wrap这个错误
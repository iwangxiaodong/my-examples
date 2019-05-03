package our.application;

import com.atomikos.jdbc.AtomikosDataSourceBean;
import java.util.Properties;
import org.apache.ibatis.datasource.DataSourceFactory;
import javax.sql.DataSource;

public class JtaDataSourceFactory implements DataSourceFactory {

    AtomikosDataSourceBean ds;

    public JtaDataSourceFactory() {
        System.out.println("JtaDataSourceFactory()");
        this.ds = new AtomikosDataSourceBean();
    }

    @Override
    public void setProperties(Properties props) {
        this.ds.setUniqueResourceName(JtaDataSourceFactory.class.getName());
        this.ds.setXaDataSourceClassName("org.mariadb.jdbc.MariaDbDataSource");
        this.ds.setMaxPoolSize(30); //  可选    

        //  maxLifetime已取代testQuery - https://www.atomikos.com/Blog/ExtremeTransactions3dot9dot0
        this.ds.setMaxLifetime(5 * 60);    //  连接最大生命周期，默认0秒无限制；若遇到Connection is closed，则会在最多maxLifetime间隔后重连。
        System.out.println("Atomikos maxLifetime - " + this.ds.getMaxLifetime());

        this.ds.setXaProperties(props);
    }

    @Override
    public DataSource getDataSource() {
        return this.ds;
    }
}

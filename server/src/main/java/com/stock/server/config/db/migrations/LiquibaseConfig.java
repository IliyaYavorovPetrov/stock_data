package com.stock.server.config.db.migrations;

import jakarta.annotation.PostConstruct;
import liquibase.Contexts;
import liquibase.LabelExpression;
import liquibase.Liquibase;
import liquibase.database.jvm.JdbcConnection;
import liquibase.exception.LiquibaseException;
import liquibase.resource.ClassLoaderResourceAccessor;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.jdbc.DataSourceBuilder;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import javax.sql.DataSource;
import java.sql.DatabaseMetaData;
import java.sql.SQLException;

@Configuration
public class LiquibaseConfig {
    public DataSource dataSourceMariadb;
    public DataSource dataSourceCassandra;

    public LiquibaseConfig(@Qualifier("mariadb") DataSource dataSourceMariadb, @Qualifier("cassandra") DataSource dataSourceCassandra) {
        this.dataSourceMariadb = dataSourceMariadb;
        this.dataSourceCassandra = dataSourceCassandra;
    }

    @PostConstruct
    public void runMigrations() throws SQLException, LiquibaseException {
        Liquibase cassandraLiquibase = new Liquibase("classpath:cassandra-changelog.xml", new ClassLoaderResourceAccessor(), new JdbcConnection(dataSourceCassandra.getConnection()));
        cassandraLiquibase.update(new Contexts(), new LabelExpression());

        Liquibase mariaDBLiquibase = new Liquibase("classpath:mariadb-changelog.xml", new ClassLoaderResourceAccessor(), new JdbcConnection(dataSourceMariadb.getConnection()));
        mariaDBLiquibase.update(new Contexts(), new LabelExpression());
    }
}

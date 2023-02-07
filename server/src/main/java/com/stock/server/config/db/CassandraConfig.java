package com.stock.server.config.db;

import com.datastax.oss.driver.api.core.CqlSession;
import org.springframework.boot.jdbc.DataSourceBuilder;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.jdbc.datasource.SingleConnectionDataSource;

import javax.sql.DataSource;

@Configuration
public class CassandraConfig {
    @Bean(name="cassandra")
    public DataSource dataSourceCassandra() {
        return DataSourceBuilder
                .create()
                .type(SingleConnectionDataSource.class)
                .build();
    }
    @Bean
    public CqlSession toCqlSession() {
        return CqlSession.builder().build();
    }
}

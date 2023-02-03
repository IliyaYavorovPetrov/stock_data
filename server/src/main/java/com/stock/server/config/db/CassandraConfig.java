package com.stock.server.config.db;

import com.datastax.oss.driver.api.core.CqlSession;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Primary;
import org.springframework.data.cassandra.config.DefaultCqlBeanNames;

@Configuration
public class CassandraConfig {
    @Bean
    public CqlSession toCqlSession() {
        return CqlSession.builder().build();
    }
}
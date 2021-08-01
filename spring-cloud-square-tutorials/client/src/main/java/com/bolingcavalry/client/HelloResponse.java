package com.bolingcavalry.client;

import lombok.AllArgsConstructor;
import lombok.Data;

import java.util.Date;

/**
 * @author zq2599@gmail.com
 * @Title: 普通bean
 * @Package
 * @Description:
 * @date 8/1/21 7:37 上午
 */
@Data
@AllArgsConstructor
public class HelloResponse {

    private String name;

    private Date date;

    private String description;
}
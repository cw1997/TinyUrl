-- tinyurl.sql
/*******************************************
 * TinyUrl Ver 1.0.0 短网址服务
 * author:昌维 [github.com/cw1997/TinyUrl]
 * email:867597730@qq.com
 *******************************************/

-- 核心数据表文件，请勿随意改动。
-- 建议不要随意修改线上环境的short_url短网址长度，
-- 如果确实要修改，请手动修改数据表中的shorturl字段，或者删表重新生成。

CREATE TABLE IF NOT EXISTS `_PREFIX_url` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `longurl` varchar(7713) NOT NULL COMMENT '各浏览器HTTP Get请求URL最大长度并不相同，几类常用浏览器最大长度及超过最大长度后提交情况如下：\r\n\r\nIE6.0                :url最大长度2083个字符，超过最大长度后无法提交。\r\nIE7.0                :url最大长度2083个字符，超过最大长度后仍然能提交，但是只能传过去2083个字符。\r\nfirefox 3.0.3     :url最大长度7764个字符，超过最大长度后无法提交。\r\nOpera 9.52       :url最大长度7648个字符，超过最大长度后无法提交。\r\nGoogle Chrome 2.0.168   :url最大长度7713个字符，超过最大长度后无法提交\r\n\r\n且7713刚好为3的整数倍（一个utf-8编码字符占用三个字节）',
  `shorturl` char(_LENGTH_) NOT NULL,
  `add_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `add_ip` varchar(21) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `shorturl` (`shorturl`) USING HASH,
  KEY `longurl` (`longurl`(255))
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
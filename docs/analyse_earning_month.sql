
DROP PROCEDURE IF EXISTS analyse_earning_month;
DELIMITER //
-- description 支出统计
-- author lujiang
-- created_at 2021-08-01
create procedure analyse_earning_month()
begin
    DECLARE curYear integer;
    DECLARE curMonth integer;

    set curYear = date_format(curdate(),'%Y');
    set curMonth =date_format(curdate(),'%m');

    SET SQL_SAFE_UPDATES = 0;

    -- 按【user_id, y, star_date,end_date, category_id】 进行统计
    delete from wealth_earning_statistics_months where year = curYear and month = curMonth;
    
    insert into wealth_earning_statistics_months(`year`, `month`, user_id, total)
    select 
        curYear,
        curMonth,
        user_id,
        sum(amount) total
    from wealth_earnings
    group by user_id;

end
//
DELIMITER ;
